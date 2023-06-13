package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"x/httpserver/internal/db"
)

func NewGetUser(connection *sql.DB) *getUser {
	return &getUser{connection: connection}
}

type getUser struct {
	connection *sql.DB
}

type getUserRequest struct {
	Id int
}

func (h *getUser) Handle(writer http.ResponseWriter, request *http.Request) {
	err := h.validateRequest(request)
	if err != nil {
		SendError(writer, http.StatusBadRequest, err)
		return
	}

	data, err := h.readData(request)
	if err != nil {
		SendError(writer, http.StatusBadRequest, err)
		return
	}

	err = h.validateData(data)
	if err != nil {
		SendError(writer, http.StatusBadRequest, err)
		return
	}

	users, err := db.GetUser(h.connection, data.Id)
	if err != nil {
		SendError(writer, http.StatusInternalServerError, err)
		return
	}

	usersResult, err := json.Marshal(users)
	if err != nil {
		SendError(writer, http.StatusInternalServerError, err)
		return
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(usersResult)
}

func (h *getUser) validateRequest(request *http.Request) error {
	if request.Method != "GET" {
		return errors.New("not found")
	}
	if len(request.Header["Content-Type"]) == 0 || request.Header["Content-Type"][0] != "application/json" {
		return errors.New("request must be in JSON format (Content-Type: application/json)")
	}
	return nil
}

func (h *getUser) validateData(data *getUserRequest) error {
	if data.Id < 0 {
		return errors.New("invalid user Id")
	}

	return nil
}

func (h *getUser) readData(request *http.Request) (*getUserRequest, error) {
	body, _ := io.ReadAll(request.Body)
	defer request.Body.Close()

	var data *getUserRequest
	err := json.Unmarshal(body, &data)

	return data, err
}

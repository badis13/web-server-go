package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"x/httpserver/internal/db"
)

func NewGetUsers(connection *sql.DB) *getUsers {
	return &getUsers{connection: connection}
}

type getUsers struct {
	connection *sql.DB
}

type getUsersRequest struct {
	Limit  int
	Offset int
}

func (h *getUsers) Handle(writer http.ResponseWriter, request *http.Request) {
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

	users, err := db.GetUsers(h.connection, data.Limit, data.Offset)
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

func (h *getUsers) validateRequest(request *http.Request) error {
	if request.Method != "GET" {
		return errors.New("not found")
	}
	if len(request.Header["Content-Type"]) == 0 || request.Header["Content-Type"][0] != "application/json" {
		return errors.New("request must be in JSON format (Content-Type: application/json)")
	}
	return nil
}

func (h *getUsers) validateData(data *getUsersRequest) error {
	if data.Limit < 1 || data.Limit > 100 {
		return errors.New("limit must be in range [1;100]")
	}
	if data.Offset < 0 {
		return errors.New("offset must be more then zero")
	}

	return nil
}

func (h *getUsers) readData(request *http.Request) (*getUsersRequest, error) {
	body, _ := io.ReadAll(request.Body)
	defer request.Body.Close()

	var data *getUsersRequest
	err := json.Unmarshal(body, &data)

	return data, err
}

package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"x/httpserver/internal/db"
)

func NewUpdateUser(connection *sql.DB) *updateUser {
	return &updateUser{connection: connection}
}

type updateUser struct {
	connection *sql.DB
}

type updateUserRequest struct {
	Id  int
	Age int
}

func (h *updateUser) Handle(writer http.ResponseWriter, request *http.Request) {
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

	users, err := db.UpdateUser(h.connection, data.Id, data.Age)
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

func (h *updateUser) validateRequest(request *http.Request) error {
	if request.Method != "PATCH" {
		return errors.New("not found")
	}
	if len(request.Header["Content-Type"]) == 0 || request.Header["Content-Type"][0] != "application/json" {
		return errors.New("request must be in JSON format (Content-Type: application/json)")
	}
	return nil
}

func (h *updateUser) validateData(data *updateUserRequest) error {
	if data.Id < 0 {
		return errors.New("invalid Id")
	}
	if data.Age < 18 {
		return errors.New("invalid Age")
	}

	return nil
}

func (h *updateUser) readData(request *http.Request) (*updateUserRequest, error) {
	body, _ := io.ReadAll(request.Body)
	defer request.Body.Close()

	var data *updateUserRequest
	err := json.Unmarshal(body, &data)

	return data, err
}

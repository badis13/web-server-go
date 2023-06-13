package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"x/httpserver/internal/db"
)

func NewDeleteUser(connection *sql.DB) *deleteUser {
	return &deleteUser{connection: connection}
}

type deleteUser struct {
	connection *sql.DB
}

type deleteUserRequest struct {
	Id int
}

func (h *deleteUser) Handle(writer http.ResponseWriter, request *http.Request) {
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

	users, err := db.DeleteUser(h.connection, data.Id)
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

func (h *deleteUser) validateRequest(request *http.Request) error {
	if request.Method != "DELETE" {
		return errors.New("not found")
	}
	if len(request.Header["Content-Type"]) == 0 || request.Header["Content-Type"][0] != "application/json" {
		return errors.New("request must be in JSON format (Content-Type: application/json)")
	}
	return nil
}

func (h *deleteUser) validateData(data *deleteUserRequest) error {
	if data.Id < 0 {
		return errors.New("invalid Id")
	}

	return nil
}

func (h *deleteUser) readData(request *http.Request) (*deleteUserRequest, error) {
	body, _ := io.ReadAll(request.Body)
	defer request.Body.Close()

	var data *deleteUserRequest
	err := json.Unmarshal(body, &data)

	return data, err
}

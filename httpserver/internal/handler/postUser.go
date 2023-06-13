package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"x/httpserver/internal/db"
)

func NewPostUser(connection *sql.DB) *postUser {
	return &postUser{connection: connection}
}

type postUser struct {
	connection *sql.DB
}

type postUserRequest struct {
	Id   int
	Name string
	Age  int
}

func (h *postUser) Handle(writer http.ResponseWriter, request *http.Request) {
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

	users, err := db.PostUser(h.connection, data.Id, data.Name, data.Age)
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

func (h *postUser) validateRequest(request *http.Request) error {
	if request.Method != "POST" {
		return errors.New("not found")
	}
	if len(request.Header["Content-Type"]) == 0 || request.Header["Content-Type"][0] != "application/json" {
		return errors.New("request must be in JSON format (Content-Type: application/json)")
	}
	return nil
}

func (h *postUser) validateData(data *postUserRequest) error {
	if data.Name == "" || data.Age < 18 {
		return errors.New("invalid data")
	}

	return nil
}

func (h *postUser) readData(request *http.Request) (*postUserRequest, error) {
	body, _ := io.ReadAll(request.Body)
	defer request.Body.Close()

	var data *postUserRequest
	err := json.Unmarshal(body, &data)

	return data, err
}

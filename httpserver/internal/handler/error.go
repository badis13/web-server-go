package handler

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func SendError(writer http.ResponseWriter, statusCode int, err error) {
	bytes, err := json.Marshal(&errorResponse{Error: err.Error()})

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	writer.Write(bytes)
}

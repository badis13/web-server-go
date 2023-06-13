package main

import (
	"encoding/json"
	"net/http"
)

func handlerError(writer http.ResponseWriter, request *http.Request) bool {
	if request.Method != "GET" {
		sendError(writer, http.StatusNotFound, "not found")
		return true
	}

	if len(request.Header["Content-Type"]) == 0 || request.Header["Content-Type"][0] != "application/json" {
		sendError(writer, http.StatusBadRequest, "invalid content type")
		return true
	}

	if request.Method != "GET" {
		sendError(writer, http.StatusNotFound, "not found")
		return true
	}

	if len(request.Header["Content-Type"]) == 0 || request.Header["Content-Type"][0] != "application/json" {
		sendError(writer, http.StatusBadRequest, "invalid content type")
		return true
	}
	return false

}

func sendError(writer http.ResponseWriter, statusCode int, message string) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	response, _ := json.Marshal(&ResponseError{Error: message})
	writer.Write(response)
}

package main

import (
	"net/http"
)

func about(writer http.ResponseWriter, request *http.Request) {
	if handlerError(writer, request) {
		return
	}

	response := []byte("I am a student with no experience as a programmer")

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

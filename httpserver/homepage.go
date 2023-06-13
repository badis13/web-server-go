package main

import (
	"net/http"
)

func homepage(writer http.ResponseWriter, request *http.Request) {
	if handlerError(writer, request) {
		return
	}

	response := []byte("This site is a learning project")

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
)

func GetAllUsers(connection *sql.DB) {

	body, err := io.ReadAll(context.Request.Body)
	defer context.Request.Body.Close()

	if err != nil {
		sendError(context.Writer, http.StatusBadRequest, "can't read body")
		return
	}

	var req *RequestHomepage
	err = json.Unmarshal(body, &req)
	if err != nil {
		sendError(context.Writer, http.StatusBadRequest, "invalid json")
		return
	}

	if req.Limit < 1 || req.Limit > 100 {
		sendError(context.Writer, http.StatusBadRequest, "limit must be in range [1;100]")
		return
	}
	users := dbParse(c)
	if req.Limit > len(users) {
		req.Limit = len(users)
	}

	response, _ := json.Marshal(users[:req.Limit])

	context.Writer.Header().Add("Content-Type", "application/json")
	context.Writer.WriteHeader(http.StatusOK)
	context.Writer.Write(response)
}

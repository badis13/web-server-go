package main

import (
	"net/http"
	"x/httpserver/internal/db"
	"x/httpserver/internal/handler"

	_ "github.com/lib/pq"
)

func main() {
	connection, err := db.GetConnection()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler.NewGetUsers(connection).Handle)
	http.HandleFunc("/adduser/", handler.NewPostUser(connection).Handle)
	http.HandleFunc("/getuserid/", handler.NewGetUser(connection).Handle)
	http.HandleFunc("/deluser/", handler.NewDeleteUser(connection).Handle)
	http.HandleFunc("/updateuser/", handler.NewUpdateUser(connection).Handle)

	http.ListenAndServe(":3333", nil)
}

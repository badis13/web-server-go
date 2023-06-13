package handler

import (
	"database/sql"
)

func dbParse(connection *sql.DB) []ResponseUser {

	var users []ResponseUser

	user, err := connection.Query("select id, name, age from human")
	if err != nil {
		panic(err)
	}
	defer user.Close()
	for user.Next() {
		p := ResponseUser{}
		err := user.Scan(&p.Id, &p.Name, &p.Age)
		if err != nil {
			println(err.Error())
			continue
		}
		users = append(users, p)
	}
	return users
}

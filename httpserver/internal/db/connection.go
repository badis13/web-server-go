package db

import (
	"database/sql"
	"fmt"
)

func GetConnection() (*sql.DB, error) {
	const (
		host     = "localhost"
		port     = 5432
		user     = "evgen"
		password = "321654"
		dbname   = "test2db"
		sslmode  = "disable"
	)

	connDb := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)
	return sql.Open("postgres", connDb)
}

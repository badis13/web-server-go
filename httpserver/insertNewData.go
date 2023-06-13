package main

import (
	"fmt"
)

func insertnewdata(req *ResponseUser) (int64, error) {

	result, err := Сonnection.Exec("INSERT INTO human (id, name, age) VALUES (?, ?, ?)", req.Id, req.Name, req.Age)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

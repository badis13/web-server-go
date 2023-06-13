package db

import "database/sql"

type user struct {
	Id   int
	Name string
	Age  int
}

func GetUsers(connection *sql.DB, limit int, offset int) ([]user, error) {
	users, err := connection.Query("SELECT id, name, age FROM human LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	defer users.Close()

	result := make([]user, 0, 10)
	for users.Next() {
		row := user{}
		err := users.Scan(&row.Id, &row.Name, &row.Age)
		if err != nil {
			return nil, err
		}

		result = append(result, row)
	}

	return result, err
}

func GetUser(connection *sql.DB, id int) ([]user, error) {
	userFind := connection.QueryRow("SELECT id, name, age FROM human WHERE id = $1", id)
	result := make([]user, 0, 1)
	row := user{}
	err1 := userFind.Scan(&row.Id, &row.Name, &row.Age)
	if err1 != nil {
		return nil, err1
	}

	result = append(result, row)

	return result, nil

}

func PostUser(connection *sql.DB, id int, name string, age int) (int64, error) {
	userAdd, err := connection.Exec("INSERT INTO human (id, name, age) values ($1, $2, $3)", id, name, age)
	userId, err1 := userAdd.RowsAffected()
	if err != nil {
		return -1, err
	}
	if err1 != nil {
		return -1, err
	}

	return userId, nil

}

func DeleteUser(connection *sql.DB, id int) (int64, error) {
	userDel, err := connection.Exec("DELETE FROM human WHERE id = $1", id)
	numDel, err1 := userDel.RowsAffected()
	if err != nil {
		return -1, err
	}
	if err1 != nil {
		return -1, err
	}

	return numDel, nil

}

func UpdateUser(connection *sql.DB, id int, age int) (int64, error) {
	userUpdate, err := connection.Exec("UPDATE human set age = $1 WHERE id = $2", age, id)
	numDel, err1 := userUpdate.RowsAffected()
	if err != nil {
		return -1, err
	}
	if err1 != nil {
		return -1, err
	}

	return numDel, nil
}

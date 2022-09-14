package models

import "app/db"

func Update(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	//resultado
	res, err := conn.Exec(`UPDATE todos SET title=$2, description=$3, dibe=$4 WHERE id=$1`, id, todo.Title, todo.Description, todo.Done)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
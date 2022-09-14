package db

import (
	"app/configs"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error){
	conf := configs.GetDB()
	
	//string connection
	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	//connection
	conn, err := sql.Open("postgre", sc)

	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping()

	return conn, err
}
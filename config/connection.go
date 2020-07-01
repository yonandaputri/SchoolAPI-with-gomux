package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectionDB() *sql.DB {
	connect := "root:Viontin@12@tcp(localhost:3306)/school"
	db, err := sql.Open("mysql", connect)
	if err != nil {
		log.Panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

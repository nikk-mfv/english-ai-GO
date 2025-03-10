package db

import (
	"database/sql"
)

var db *sql.DB

func Connect() error {
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/english-ai")
	if err != nil {
		return err
	}
	return err
}

func GetDB() *sql.DB {
	return db
}

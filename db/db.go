package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func OpenConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "./alura_loja.db")
	if err != nil {
		panic(err.Error())
	}
	return db
}

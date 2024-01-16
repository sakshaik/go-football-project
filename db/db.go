package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	MyDB, err := sql.Open("sqlite3", "football.db")
	if err != nil {
		panic("Could not connect to DB")
	}
	DB = MyDB
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	CreateTables()
	InsertMasterData()
}

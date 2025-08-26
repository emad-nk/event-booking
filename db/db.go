package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS user (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT not null,
		password TEXT not null
	)`

	_, err := DB.Exec(createUserTable)

	if err != nil {
		panic("Could not create user table.")
	}

	createEventTable := `
	CREATE TABLE IF NOT EXISTS event (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT not null,
		description TEXT not null,
		location TEXT not null,
		dateTime DATETIME not null,
		user_id INTEGER,
		foreign key(user_id) references user(id)
	)`

	_, err = DB.Exec(createEventTable)

	if err != nil {
		panic("Could not create event table.")
	}
}

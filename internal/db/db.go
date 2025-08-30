package db

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var ErrDBConnectionFailed = errors.New("failed to connect to database")
var ErrDBPingFailed = errors.New("failed to ping database")

func MustConnect(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(ErrDBConnectionFailed)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(ErrDBPingFailed)
	}
	return db
}
package db

import (
	"database/sql"
	"embed"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

var ErrDBConnectionFailed = errors.New("failed to connect to database")
var ErrDBPingFailed = errors.New("failed to ping database")

func MustConnect() *sql.DB {
	dbPath := "./box.db"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(ErrDBConnectionFailed)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(ErrDBPingFailed)
	}
	return db
}

//go:embed migrations/*.sql
var embedMigrations embed.FS

func MustMigrate(db *sql.DB) {
  goose.SetBaseFS(embedMigrations)	
  // Do not log migrations
  goose.SetLogger(goose.NopLogger())
  
  if err := goose.SetDialect("sqlite3"); err != nil {
    log.Fatal(err)
  }
  if err := goose.Up(db, "migrations"); err != nil {
    log.Fatal(err)
  }
}

func EnsureDB() *sql.DB {
	db := MustConnect()
	MustMigrate(db)

	return db
}
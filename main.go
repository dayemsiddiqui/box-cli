package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dayemsiddiqui/box-cli/internal/db"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	fmt.Println("Hello, World!")

	dbURL := os.Getenv("DB_URL")
	db := db.MustConnect(dbURL)
	defer db.Close()

	// Migrate the database columns

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS tasks (id TEXT PRIMARY KEY, name TEXT, description TEXT, created_at TIMESTAMP, updated_at TIMESTAMP);
	`
	if _, err := db.Exec(sqlStmt); err != nil {
		log.Fatal(err)
	}
	
}
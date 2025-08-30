package main

import (
	"fmt"
	"os"

	"github.com/dayemsiddiqui/box-cli/internal/db"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	fmt.Println("Hello, World!")

	dbURL := os.Getenv("DB_URL")
	db := db.MustConnect(dbURL)
	defer db.Close()

	
}
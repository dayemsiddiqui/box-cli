/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/dayemsiddiqui/box-cli/cmd"
	"github.com/dayemsiddiqui/box-cli/internal/db"
)

func main() {

	dbURL := os.Getenv("DB_URL")
	database := db.MustConnect(dbURL)
	defer database.Close()
	
	
	
	cmd.Execute()
}

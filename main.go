/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/dayemsiddiqui/box-cli/cmd"
	"github.com/dayemsiddiqui/box-cli/internal/db"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	database := db.EnsureDB()
	defer database.Close()
	
	
	
	cmd.Execute()
}

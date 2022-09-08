package main

import (
	"DeployX/cmd"
	"DeployX/database"
)

func main() {
	initializeDatabase()
	cmd.Execute()
}

func initializeDatabase() {
	database.Initialize("deployx.db")
	models.InitializeConfig()
}

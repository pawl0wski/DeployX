package main

import (
	"DeployX/cmd"
	"DeployX/database"
	"DeployX/model"
)

func main() {
	initializeDatabase()
	cmd.Execute()
}

func initializeDatabase() {
	database.Initialize("deployx.db")
	model.InitializeConfig()
}

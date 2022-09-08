package main

import (
	"DeployX/cmd"
	"DeployX/database"
	"DeployX/models"
)

func main() {
	initializeDatabaseAndModels()
	cmd.Execute()
}

func initializeDatabaseAndModels() {
	database.Initialize("deployx.db")
	models.InitializeConfig()
}

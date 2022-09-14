package main

import (
	"github.com/pawl0wski/DeployX/cmd"
	"github.com/pawl0wski/DeployX/database"
	"github.com/pawl0wski/DeployX/models"
)

func main() {
	initializeDatabaseAndModels()
	cmd.Execute()
}

func initializeDatabaseAndModels() {
	database.Initialize("deployx.db")
	models.InitializeConfig()
	models.InitializeScript()
	models.InitializeDeploymentConfiguration()
	models.InitializeProject()
}

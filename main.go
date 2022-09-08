package main

import (
	"DeployX/cmd"
	"DeployX/database"
)

func main() {
	database.Initialize("deployx.db")
	cmd.Execute()
}

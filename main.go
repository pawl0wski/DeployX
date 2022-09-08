package main

import (
	"DeployX/cmd"
	"DeployX/database"
)

func main() {
	database.Database("deployx.db")
	cmd.Execute()
}

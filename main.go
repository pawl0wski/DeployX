package main

import (
	"DeployX/cmd"
	"DeployX/initialize"
)

func main() {
	initialize.Database("deployx.db")
	cmd.Execute()
}

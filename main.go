package main

import (
	"DeployX/cmd"
	"DeployX/database"
	"DeployX/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	initializeDatabase()
	cmd.Execute()
}

func initializeDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("deployx.db"), &gorm.Config{})
	if err != nil {
		panic("Can't open database")
	}

	err = database.DBConn.AutoMigrate(&models.Config{})
	if err != nil {
		panic("Can't migrate Config")
	}
}

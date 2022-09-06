package main

import (
	"DeployX/cmd"
	"DeployX/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initializeDatabase() {
	db, err := gorm.Open(sqlite.Open("deployx.db"), &gorm.Config{})
	if err != nil {
		panic("Can't open database")
	}

	err = db.AutoMigrate(&models.Config{})
	if err != nil {
		panic("Can't migrate Config")
	}
}

func main() {
	initializeDatabase()
	cmd.Execute()
}

package database

import (
	"DeployX/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Database(databasePath string) {
	if databasePath == "" {
		databasePath = ":memory:"
	}
	var err error
	DBConn, err = gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		panic("Can't open database")
	}

	err = DBConn.AutoMigrate(&model.Config{})
	if err != nil {
		panic("Can't migrate CreateConfig")
	}
}

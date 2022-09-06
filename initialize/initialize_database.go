package initialize

import (
	"DeployX/database"
	"DeployX/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Database(databasePath string) {
	if databasePath == "" {
		databasePath = ":memory:"
	}
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		panic("Can't open database")
	}

	err = database.DBConn.AutoMigrate(&model.Config{})
	if err != nil {
		panic("Can't migrate CreateConfig")
	}
}

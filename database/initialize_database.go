package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Initialize(databasePath string) {
	if databasePath == "" {
		databasePath = ":memory:"
	}
	debugMode := getDebugModeFromDatabase(databasePath)
	var err error
	DBConn, err = gorm.Open(sqlite.Open(databasePath), &gorm.Config{Logger: getLoggerBasedOnDebugMode(debugMode)})
	if err != nil {
		panic("Can't open database")
	}
}

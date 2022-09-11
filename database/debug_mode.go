package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getDebugModeFromDatabase(databasePath string) bool {
	var err error
	DBConn, err = gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		panic("Unable to open database for debug mode information")
	}
	var result bool
	DBConn.Raw("SELECT debug_mode FROM configs LIMIT 1;").Scan(&result)
	return result
}

func getLoggerBasedOnDebugMode(debugMode bool) logger.Interface {
	if debugMode {
		return logger.Default.LogMode(logger.Silent)
	}
	return logger.Default
}

package models

import (
	"github.com/pawl0wski/DeployX/database"
	"gorm.io/gorm"
)

type Config struct {
	gorm.Model
	TextEditor string `json:"string_editor"`
	ServerPort uint16 `json:"server_port"`
	DebugMode  bool   `json:"debug_mode"`
}

// Save saves config in database
func (config *Config) Save() {
	database.DBConn.Save(config)
}

// GetFromDatabaseOrCreate update Config struct with data from database or create new Config in database
func (config *Config) GetFromDatabaseOrCreate() {
	database.DBConn.FirstOrCreate(config)
}

// DoesConfigExist return true if config exist in database
func DoesConfigExist() bool {
	var howMuchConfigs int64
	database.DBConn.Table("configs").Count(&howMuchConfigs)
	return howMuchConfigs > 0
}

func InitializeConfig() {
	database.InitializeModel(Config{})
}

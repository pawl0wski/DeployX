package models

import (
	"DeployX/database"
	"gorm.io/gorm"
)

type Config struct {
	gorm.Model
	TextEditor string `json:"string_editor"`
	ServerPort uint16 `json:"server_port"`
}

func (config *Config) Save() {
	database.DBConn.Updates(config)
}

func (config *Config) GetFromDatabaseOrCreate() {
	database.DBConn.FirstOrCreate(config)
}

func init() {
	database.InitializeModel(Config{})
}

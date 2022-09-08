package model

import (
	"DeployX/database"
	"gorm.io/gorm"
)

type Config struct {
	gorm.Model
	TextEditor string `json:"string_editor"`
	ServerPort uint   `json:"server_port"`
}

func (config *Config) Save() {
	database.DBConn.Updates(config)
}

func (config *Config) GetFromDatabaseOrCreate() {
	database.DBConn.FirstOrCreate(config)
}

func InitializeConfig() {
	database.InitializeModel(Config{})
}

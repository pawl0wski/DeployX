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

func (config *Config) Update() {
	database.DBConn.Updates(config)
}

// Refresh Config with values from database
func (config *Config) Refresh() {
	database.DBConn.FirstOrCreate(config)
}

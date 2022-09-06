package model

import (
	"gorm.io/gorm"
)

type Config struct {
	gorm.Model
	TextEditor string `json:"string_editor"`
	ServerPort uint   `json:"server_port"`
}

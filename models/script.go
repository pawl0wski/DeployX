package models

import (
	"DeployX/database"
	"gorm.io/gorm"
)

type Script struct {
	gorm.Model
	Content string `json:"content"`
}

func InitializeScript() {
	database.InitializeModel(Script{})
}

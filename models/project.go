package models

import (
	"DeployX/database"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name     string `json:"name"`
	Path     string `json:"path"`
	Password string `json:"password"`
}

func InitializeProject() {
	database.InitializeModel(Project{})
}

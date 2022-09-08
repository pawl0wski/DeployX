package models

import (
	"DeployX/database"
	"crypto/sha256"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name     string `json:"name"`
	Path     string `json:"path"`
	Password string `json:"password"`
}

func (project *Project) SetPassword(password string) {
	hasher := sha256.New()
	hashedPassword := hasher.Sum([]byte(password))
	project.Password = string(hashedPassword)
}

func InitializeProject() {
	database.InitializeModel(Project{})
}

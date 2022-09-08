package models

import (
	"DeployX/database"
	"DeployX/modules"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name     string `json:"name"`
	Path     string `json:"path"`
	Password string `json:"password"`
}

func (project *Project) SetPassword(password string) {
	project.Password = modules.HashPassword(password)
}

func (project *Project) ValidatePassword(password string) bool {
	return project.Password == modules.HashPassword(password)
}

func GetAllProjects() []Project {
	var projects []Project
	database.DBConn.Find(&projects)
	return projects
}

func InitializeProject() {
	database.InitializeModel(Project{})
}

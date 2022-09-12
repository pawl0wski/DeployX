package models

import (
	"github.com/pawl0wski/DeployX/database"
	"github.com/pawl0wski/DeployX/hasher"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name                 string `json:"name"`
	Path                 string `json:"path"`
	Password             string `json:"password"`
	BeforeDeployScript   Script `json:"before_deploy_script" gorm:"foreignKey:BeforeDeployScriptID"`
	BeforeDeployScriptID int    `json:"before_deploy_script_id"`
	AfterDeployScript    Script `json:"after_deploy_script" gorm:"foreignKey:AfterDeployScriptID"`
	AfterDeployScriptID  int    `json:"after_deploy_script_id"`
}

func (project *Project) Save() {
	database.DBConn.Save(project)
}

func (project *Project) Create() {
	database.DBConn.Create(project)
}

func (project *Project) Delete() {
	database.DBConn.Delete(project)
}

func (project *Project) SetPassword(password string) {
	project.Password = hasher.HashPassword(password)
}

func (project *Project) ValidatePassword(password string) bool {
	return project.Password == hasher.HashPassword(password)
}

func GetAllProjects() []Project {
	var projects []Project
	database.DBConn.Find(&projects)
	return projects
}

func InitializeProject() {
	database.InitializeModel(Project{})
}

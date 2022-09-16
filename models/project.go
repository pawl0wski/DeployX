package models

import (
	"errors"
	"github.com/pawl0wski/DeployX/database"
	"github.com/pawl0wski/DeployX/hasher"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name                      string                  `json:"name"`
	Path                      string                  `json:"path"`
	Password                  string                  `json:"password"`
	BeforeDeployScript        Script                  `json:"before_deploy_script" gorm:"foreignKey:BeforeDeployScriptID"`
	BeforeDeployScriptID      int                     `json:"before_deploy_script_id"`
	AfterDeployScript         Script                  `json:"after_deploy_script" gorm:"foreignKey:AfterDeployScriptID"`
	AfterDeployScriptID       int                     `json:"after_deploy_script_id"`
	DeploymentConfiguration   DeploymentConfiguration `json:"deployment_configuration" gorm:"foreignKey:DeploymentConfigurationID"`
	DeploymentConfigurationID int                     `json:"deployment_configuration_id"`
	ProjectSnapshots          []ProjectSnapshot       `json:"project_snapshots"`
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
	db := database.DBConn.Model(&Project{})
	preloadProjectAssociations(db)
	db.Find(&projects)
	return projects

}

func GetProjectByID(id int) (Project, error) {
	preloadProjectAssociations(database.DBConn)
	var project Project
	database.DBConn.Find(&project)
	projectIDUint := uint(id)
	if project.ID != projectIDUint {
		return project, errors.New("there is no project with the given id")
	}
	return project, nil
}

func preloadProjectAssociations(db *gorm.DB) {
	db.Preload("BeforeDeployScript")
	db.Preload("AfterDeployScript")
	db.Preload("DeploymentConfiguration")
}

func InitializeProject() {
	database.InitializeModel(Project{})
}

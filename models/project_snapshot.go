package models

import (
	"github.com/pawl0wski/DeployX/database"
	"gorm.io/gorm"
)

type ProjectSnapshot struct {
	gorm.Model
	Version         int     `json:"version"`
	Data            []byte  `json:"data"`
	CurrentSnapshot bool    `json:"current_snapshot"`
	Project         Project `json:"project"`
	ProjectID       int     `json:"project_id"`
}

func InitializeProjectSnapshot() {
	database.InitializeModel(ProjectSnapshot{})
}

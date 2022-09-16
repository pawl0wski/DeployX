package models

import (
	"github.com/pawl0wski/DeployX/database"
	"gorm.io/gorm"
	"strconv"
)

type ProjectSnapshot struct {
	gorm.Model
	Version         int     `json:"version"`
	Data            []byte  `json:"data"`
	CurrentSnapshot bool    `json:"current_snapshot"`
	Project         Project `json:"project"`
	ProjectID       int     `json:"project_id"`
}

func (snapshot *ProjectSnapshot) Save() {
	database.DBConn.Save(snapshot)
}

func GetLastSnapshotVersionByProjectID(id int) int {
	var snapshot ProjectSnapshot
	database.DBConn.Where("project_id = ?", []string{strconv.Itoa(int(id))}).Order("version desc").First(&snapshot)
	return snapshot.Version
}

func InitializeProjectSnapshot() {
	database.InitializeModel(ProjectSnapshot{})
}

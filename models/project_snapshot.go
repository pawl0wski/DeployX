package models

import (
	"encoding/base64"
	"errors"
	"github.com/pawl0wski/DeployX/database"
	"gorm.io/gorm"
	"strconv"
)

type ProjectSnapshot struct {
	gorm.Model
	Version         int     `json:"version"`
	Data            string  `json:"data"`
	CurrentSnapshot bool    `json:"current_snapshot"`
	Project         Project `json:"project"`
	ProjectID       int     `json:"project_id"`
}

func (snapshot *ProjectSnapshot) Save() {
	database.DBConn.Save(snapshot)
}

func (snapshot *ProjectSnapshot) ConvertDataToBinary() ([]byte, error) {
	encoder := base64.Encoding{}
	binaryData, err := encoder.DecodeString(snapshot.Data)
	if err != nil {
		return nil, errors.New("can't covert data to binary")
	}
	return binaryData, nil
}

func GetLastSnapshotVersionByProjectID(id int) int {
	var snapshot ProjectSnapshot
	database.DBConn.Where("project_id = ?", []string{strconv.Itoa(int(id))}).Order("version desc").First(&snapshot)
	return snapshot.Version
}

func InitializeProjectSnapshot() {
	database.InitializeModel(ProjectSnapshot{})
}

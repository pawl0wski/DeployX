package models

import (
	"github.com/pawl0wski/DeployX/database"
	"gorm.io/gorm"
)

type Script struct {
	gorm.Model
	Content string `json:"content"`
}

func (script *Script) Save() {
	database.DBConn.Save(script)
}

func InitializeScript() {
	database.InitializeModel(Script{})
}

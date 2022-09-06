package repository

import (
	"DeployX/database"
	"DeployX/model"
)

func UpdateConfig(config *model.Config) {
	database.DBConn.Updates(config)
}

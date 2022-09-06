package database

import (
	"DeployX/model"
)

func UpdateConfig(config *model.Config) {
	DBConn.Updates(config)
}

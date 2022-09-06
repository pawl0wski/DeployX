package repository

import (
	"DeployX/database"
	"DeployX/model"
)

func GetConfig() *model.Config {
	config := model.Config{}
	database.DBConn.First(&config)
	return &config
}

package database

import (
	"DeployX/model"
)

func GetConfig() *model.Config {
	config := model.Config{}
	DBConn.First(&config)
	return &config
}

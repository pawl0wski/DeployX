package editors

import (
	"DeployX/models"
	"DeployX/prompts"
)

func EditConfig(config *models.Config) {
	config.TextEditor = prompts.SelectTextEditor()
	config.ServerPort = prompts.GetServerPort(config.ServerPort)
	config.DebugMode = prompts.SelectDebugMode()
}

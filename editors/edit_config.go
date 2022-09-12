package editors

import (
	"github.com/pawl0wski/DeployX/models"
	"github.com/pawl0wski/DeployX/prompts"
)

func EditConfig(config *models.Config) {
	config.TextEditor = prompts.SelectTextEditor()
	config.ServerPort = prompts.GetServerPort(config.ServerPort)
	config.DebugMode = prompts.SelectDebugMode()
}

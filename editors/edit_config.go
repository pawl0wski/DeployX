package editors

import (
	"DeployX/models"
	"DeployX/prompts"
)

func EditConfig(config *models.Config) {
	printWithEditorColor("Config editor")
	config.TextEditor = prompts.SelectTextEditor()
	config.ServerPort = prompts.GetServerPort(config)
	config.DebugMode = prompts.SelectDebugMode()
}

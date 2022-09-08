package editors

import (
	"DeployX/editors/prompts"
	"DeployX/models"
)

func EditConfig(config *models.Config) {
	printWithEditorColor("Config editor")
	config.TextEditor = prompts.SelectTextEditor()
	config.ServerPort = prompts.GetServerPort(config)
}

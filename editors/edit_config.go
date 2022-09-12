package editors

import (
	"DeployX/models"
	prompts2 "DeployX/prompts"
)

func EditConfig(config *models.Config) {
	printWithEditorColor("Config editor")
	config.TextEditor = prompts2.SelectTextEditor()
	config.ServerPort = prompts2.GetServerPort(config)
	config.DebugMode = prompts2.SelectDebugMode()
}

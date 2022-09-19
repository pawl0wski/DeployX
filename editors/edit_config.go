package editors

import (
	"github.com/pawl0wski/DeployX/models"
	"github.com/pawl0wski/DeployX/prompts"
)

func EditConfig(config *models.Config) {
	textEditorPrompt := prompts.SelectTextEditorPrompt{}
	config.TextEditor = textEditorPrompt.Run()
	serverPortPrompt := prompts.GetServerPortPrompt{DefaultPort: config.ServerPort}
	config.ServerPort = serverPortPrompt.Run()
	debugModePrompt := prompts.SelectDebugModePrompt{}
	config.DebugMode = debugModePrompt.Run()
}

package editors

import (
	"github.com/pawl0wski/DeployX/models"
	"github.com/pawl0wski/DeployX/prompts"
)

type ConfigEditor struct {
	ConfigToEdit *models.Config
}

func (c *ConfigEditor) Run() {
	c.editTextEditor()
	c.editServerPort()
	c.editDebugMode()
}

func (c *ConfigEditor) editTextEditor() {
	textEditorPrompt := prompts.SelectTextEditorPrompt{}
	c.ConfigToEdit.TextEditor = textEditorPrompt.Run()
}

func (c *ConfigEditor) editServerPort() {
	serverPortPrompt := prompts.GetServerPortPrompt{DefaultPort: c.ConfigToEdit.ServerPort}
	c.ConfigToEdit.ServerPort = serverPortPrompt.Run()
}

func (c *ConfigEditor) editDebugMode() {
	debugModePrompt := prompts.SelectDebugModePrompt{}
	c.ConfigToEdit.DebugMode = debugModePrompt.Run()
}

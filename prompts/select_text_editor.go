package prompts

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/pawl0wski/DeployX/prompts/models"
)

type SelectTextEditorPrompt struct {
}

func (p *SelectTextEditorPrompt) Run() string {
	editorsThatAreInstalled := p.returnOnlyEditorsThatAreInstalled(p.getDefaultEditors())
	editorsAsString := p.convertEditorsToOptions(editorsThatAreInstalled)
	prompt := &survey.Select{Message: "Select your favorite text editor", Options: editorsAsString}
	var selectedOption string
	err := survey.AskOne(prompt, &selectedOption)
	if err != nil {
		panic("Can't get your favorite text editor")
	}
	selectedEditor := p.convertOptionToEditor(selectedOption, editorsThatAreInstalled)
	return selectedEditor.ToString()
}

func (p *SelectTextEditorPrompt) returnOnlyEditorsThatAreInstalled(editors []models.Command) []models.Command {
	var installedEditors []models.Command
	for _, editor := range editors {
		if editor.CheckIfExist() {
			installedEditors = append(installedEditors, editor)
		}
	}
	return installedEditors
}

func (p *SelectTextEditorPrompt) getDefaultEditors() []models.Command {
	return []models.Command{
		{Command: "vi"},
		{Command: "nano"},
		{Command: "code", DefaultArguments: []string{"-w"}},
		{Command: "gedit"},
		{Command: "kate"},
		{Command: "codium", DefaultArguments: []string{"-w"}},
		{Command: "vim"},
	}
}

func (p *SelectTextEditorPrompt) convertEditorsToOptions(commands []models.Command) []string {
	var commandsAsString []string
	for _, command := range commands {
		commandsAsString = append(commandsAsString, command.Command)
	}
	return commandsAsString
}

func (p *SelectTextEditorPrompt) convertOptionToEditor(selectedOption string, commands []models.Command) models.Command {
	var selectedCommand models.Command
	for _, command := range commands {
		if command.Command == selectedOption {
			selectedCommand = command
		}
	}
	return selectedCommand
}

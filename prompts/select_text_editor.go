package prompts

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/pawl0wski/DeployX/prompts/models"
)

func SelectTextEditor() string {
	editorsThatAreInstalled := returnOnlyEditorsThatAreInstalled(getDefaultEditors())
	editorsAsString := convertEditorsToOptions(editorsThatAreInstalled)
	prompt := &survey.Select{Message: "Select your favorite text editor", Options: editorsAsString}
	var selectedOption string
	err := survey.AskOne(prompt, &selectedOption)
	if err != nil {
		panic("Can't get your favorite text editor")
	}
	selectedEditor := convertOptionToEditor(selectedOption, editorsThatAreInstalled)
	return selectedEditor.ToString()
}

func returnOnlyEditorsThatAreInstalled(editors []models.Command) []models.Command {
	var installedEditors []models.Command
	for _, editor := range editors {
		if editor.CheckIfExist() {
			installedEditors = append(installedEditors, editor)
		}
	}
	return installedEditors
}

func getDefaultEditors() []models.Command {
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

func convertEditorsToOptions(commands []models.Command) []string {
	var commandsAsString []string
	for _, command := range commands {
		commandsAsString = append(commandsAsString, command.Command)
	}
	return commandsAsString
}

func convertOptionToEditor(selectedOption string, commands []models.Command) models.Command {
	var selectedCommand models.Command
	for _, command := range commands {
		if command.Command == selectedOption {
			selectedCommand = command
		}
	}
	return selectedCommand
}

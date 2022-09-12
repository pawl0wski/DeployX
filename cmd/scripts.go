package cmd

import (
	"DeployX/models"
	"DeployX/temp"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

// scriptsCmd represents the scripts command
var scriptsCmd = &cobra.Command{
	Use:   "scripts",
	Short: "Manage scripts",
	Run:   runScripts,
}

func runScripts(cmd *cobra.Command, args []string) {
	project := selectProject()
	script := askUserWhatScriptToEdit(project)
	editScriptUsingTextEditorDefinedInConfig(script)
	script.Save()
	project.Save()
}

func askUserWhatScriptToEdit(project *models.Project) *models.Script {
	prompt := &survey.Select{Options: []string{"Before deployment script", "After deployment script"}}
	var selection string
	err := survey.AskOne(prompt, &selection)
	if err != nil {
		panic("Can't ask user what script to edit")
	}
	var script *models.Script
	if selection == "Before deployment script" {
		return &project.BeforeDeployScript
	} else {
		return &project.AfterDeployScript
	}
	return script
}

func editScriptUsingTextEditorDefinedInConfig(script *models.Script) {
	file := temp.FileCreator(script.Content)
	config := models.Config{}
	config.GetFromDatabaseOrCreate()
	textEditorCommand := exec.Command(config.TextEditor, file.Name())
	textEditorCommand.Stdin = os.Stdin
	textEditorCommand.Stdout = os.Stdout
	err := textEditorCommand.Run()
	if err != nil {
		panic(fmt.Sprintf("Can't execute %s", config.TextEditor))
	}
	script.Content = temp.GetContentFromTempFile(file)
}

func init() {
	rootCmd.AddCommand(scriptsCmd)
}

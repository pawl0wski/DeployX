package cmd

import (
	"github.com/pawl0wski/DeployX/editors"
	"github.com/pawl0wski/DeployX/models"
	"github.com/pawl0wski/DeployX/prompts"
	"github.com/spf13/cobra"
)

// scriptsCmd represents the scripts command
var scriptsCmd = &cobra.Command{
	Use:   "scripts",
	Short: "Manage scripts",
	Run:   runScripts,
}

func runScripts(cmd *cobra.Command, args []string) {
	project := prompts.SelectProject()
	script := prompts.SelectScriptToEdit(project)
	editScriptUsingTextEditorDefinedInConfig(script)
	script.Save()
	project.Save()
}

func editScriptUsingTextEditorDefinedInConfig(script *models.Script) {
	config := models.Config{}
	config.GetFromDatabaseOrCreate()
	editors.EditScript(script, config.TextEditor)
}

func init() {
	rootCmd.AddCommand(scriptsCmd)
}

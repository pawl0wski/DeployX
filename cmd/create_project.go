package cmd

import (
	"DeployX/editors"
	"DeployX/models"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// createProjectCmd represents the createProject command
var createProjectCmd = &cobra.Command{
	Use:   "create-project",
	Short: "Create new project",
	Run:   runCreateProject,
}

func runCreateProject(cmd *cobra.Command, args []string) {
	color.Blue("Project creator")
	project := models.Project{}
	editors.EditProject(&project)
	project.Create()
}

func init() {
	rootCmd.AddCommand(createProjectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createProjectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createProjectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

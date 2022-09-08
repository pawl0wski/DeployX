package cmd

import (
	"DeployX/models"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"os"
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
	project.Name = getProjectName()
	project.Path = getPath()
	project.SetPassword(getPassword())
}

func getProjectName() string {
	prompt := promptui.Prompt{Label: "Name"}
	result, err := prompt.Run()
	if err != nil {
		panic("Can't get project's name")
	}
	return result
}

func getPath() string {
	validator := func(input string) error {
		_, err := os.Stat(input)
		if err != nil {
			return errors.New(fmt.Sprintf("path \"%s\" not exist", input))
		}
		return nil
	}
	prompt := promptui.Prompt{Label: "Path", Validate: validator}
	result, err := prompt.Run()
	if err != nil {
		panic("Can't get project's path")
	}
	return result
}

func getPassword() string {
	prompt := promptui.Prompt{Label: "Password", HideEntered: true}
	result, err := prompt.Run()
	if err != nil {
		panic("Can't get password")
	}
	return result
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

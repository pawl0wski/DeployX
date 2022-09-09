package cmd

import (
	"DeployX/models"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

// editProjectCmd represents the editProject command
var editProjectCmd = &cobra.Command{
	Use:   "edit-project [project's id to edit]",
	Short: "Edit existing project",
	Args:  cobra.MinimumNArgs(1),
	Run:   runEditProject,
}

func runEditProject(cmd *cobra.Command, args []string) {
	project := validateArgAndReturnProject(args)
	fmt.Println(project.ID)
}

func validateArgAndReturnProject(args []string) models.Project {
	firstArg := args[0]
	projectId, err := strconv.ParseUint(firstArg, 10, 32)
	if err != nil {
		color.Red("%s is not a number", firstArg)
		os.Exit(1)
	}
	project := models.GetProjectById(projectId)
	if project.ID == 0 {
		color.Red("Can't find project with id %d", projectId)
		os.Exit(1)
	}
	return project
}

func init() {
	rootCmd.AddCommand(editProjectCmd)
}

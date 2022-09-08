package cmd

import (
	"DeployX/models"
	"fmt"
	"github.com/spf13/cobra"
)

// projectsCmd represents the projects command
var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Show all projects",
	Run:   runProjects,
}

func runProjects(cmd *cobra.Command, args []string) {
	fmt.Println("Projects:")
	projects := models.GetAllProjects()
	for _, project := range projects {
		fmt.Printf("- %d: %s\n", project.ID, project.Name)
	}
}

func init() {
	rootCmd.AddCommand(projectsCmd)
}

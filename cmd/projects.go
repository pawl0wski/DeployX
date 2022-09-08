package cmd

import (
	"DeployX/models"
	"fmt"
	"github.com/fatih/color"
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
		color.Green(" %s\n", project.Name)
		fmt.Printf("  - id: %d\n", project.ID)
		fmt.Printf("  - path: %s\n", project.Path)
	}
}

func init() {
	rootCmd.AddCommand(projectsCmd)
}

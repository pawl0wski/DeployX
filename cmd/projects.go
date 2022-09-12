package cmd

import (
	"DeployX/editors"
	"DeployX/models"
	"DeployX/prompts"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// projectsCmd represents the projects command
var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Manage your projects",
	Run:   runProjects,
}

func runProjects(cmd *cobra.Command, args []string) {
	action := askTheUserForAction()
	switch action {
	case "List all projects":
		listAllProjects()
	case "Create project":
		createProject()
	case "Edit project":
		editProject()
	case "Delete project":
		deleteProject()
	}
}

func askTheUserForAction() string {
	prompt := &survey.Select{
		Message: "Choose what you want to do",
		Options: []string{"List all projects", "Create project", "Edit project", "Delete project"}}
	var selection string
	err := survey.AskOne(prompt, &selection)
	if err != nil {
		panic("Can't ask user")
	}
	return selection
}

func listAllProjects() {
	fmt.Println("Your projects:")
	projects := models.GetAllProjects()
	for _, project := range projects {
		color.Green(" %s\n", project.Name)
		fmt.Printf("  - id: %d\n", project.ID)
		fmt.Printf("  - path: %s\n", project.Path)
	}
}

func createProject() {
	project := models.Project{}
	editors.EditProject(&project)
	project.Create()
	color.Green("Your new project will have ID %d", project.ID)
}

func editProject() {
	project := prompts.SelectProject()
	editors.EditProject(project)
	project.Save()
}

func deleteProject() {
	project := prompts.SelectProject()
	prompt := &survey.Confirm{Message: fmt.Sprintf("You sure want to delete %s?", project.Name), Default: false}
	var decision bool
	err := survey.AskOne(prompt, &decision)
	if err != nil {
		panic("The deletion cannot be assured\n\n")
	}
	project.Delete()
	color.Green(fmt.Sprintf("Project %s successfuly deleted", project.Name))
}

func init() {
	rootCmd.AddCommand(projectsCmd)
}

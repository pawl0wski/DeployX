package prompts

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/pawl0wski/DeployX/models"
	"os"
)

type SelectProjectPrompt struct {
}

func (p *SelectProjectPrompt) Run() *models.Project {
	projects := models.GetAllProjects()
	if len(projects) == 0 {
		color.Red("You don't have any projects. Make one first.")
		os.Exit(0)
	}
	var promptOptions []string
	for _, project := range projects {
		promptOptions = append(promptOptions, project.Name)
	}
	prompt := &survey.Select{Message: "Pick project", Options: promptOptions}
	var selectedProjectName string
	err := survey.AskOne(prompt, &selectedProjectName)
	if err != nil {
		panic("Can't select project from user")
	}
	for _, project := range projects {
		if project.Name == selectedProjectName {
			return &project
		}
	}
	panic("Selected unknown project name")
}

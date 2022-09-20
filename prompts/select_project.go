package prompts

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/pawl0wski/DeployX/models"
	"os"
)

type SelectProjectPrompt struct {
	Projects []models.Project
}

func (p *SelectProjectPrompt) Run() *models.Project {
	p.exitIfUserDontHaveAnyProject()
	options := p.changeProjectsToOptions(p.Projects)
	prompt := p.preparePrompt(options)
	selection := p.askUser(prompt)
	project := p.changeSelectionToProject(selection)
	if project == nil {
		panic("Selected unknown project name")
	}
	return project
}

func (p *SelectProjectPrompt) exitIfUserDontHaveAnyProject() {
	if len(p.Projects) == 0 {
		color.Red("You don't have any projects. Make one first.")
		os.Exit(0)
	}
}

func (p *SelectProjectPrompt) changeProjectsToOptions(projects []models.Project) []string {
	var options []string
	for _, project := range projects {
		options = append(options, project.Name)
	}
	return options
}

func (p *SelectProjectPrompt) preparePrompt(options []string) *survey.Select {
	return &survey.Select{Message: "Pick project", Options: options}
}

func (p *SelectProjectPrompt) askUser(prompt *survey.Select) string {
	var selection string
	err := survey.AskOne(prompt, &selection)
	if err != nil {
		panic("Can't select project from user")
	}
	return selection
}

func (p *SelectProjectPrompt) changeSelectionToProject(selection string) *models.Project {
	for _, project := range p.Projects {
		if project.Name == selection {
			return &project
		}
	}
	return nil
}

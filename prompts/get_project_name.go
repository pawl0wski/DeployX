package prompts

import (
	"DeployX/models"
	"github.com/AlecAivazis/survey/v2"
)

func GetProjectName(project *models.Project) string {
	prompt := &survey.Input{Message: "Name", Default: project.Name}
	var name string
	err := survey.AskOne(prompt, &name)
	if err != nil {
		panic("Can't get project's name")
	}
	return name
}

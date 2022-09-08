package prompts

import (
	"DeployX/models"
	"github.com/manifoldco/promptui"
)

func GetProjectName(project *models.Project) string {
	prompt := promptui.Prompt{Label: "Name", Default: project.Name}
	result, err := prompt.Run()
	if err != nil {
		panic("Can't get project's name")
	}
	return result
}

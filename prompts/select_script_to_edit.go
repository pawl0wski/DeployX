package prompts

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/pawl0wski/DeployX/models"
)

func SelectScriptToEdit(project *models.Project) *models.Script {
	prompt := &survey.Select{Message: "What script do you want to edit?", Options: []string{"Before deployment script", "After deployment script"}}
	var selection string
	err := survey.AskOne(prompt, &selection)
	if err != nil {
		panic("Can't ask user what script to edit")
	}
	var script *models.Script
	if selection == "Before deployment script" {
		return &project.BeforeDeployScript
	} else {
		return &project.AfterDeployScript
	}
	return script
}

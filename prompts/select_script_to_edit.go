package prompts

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/pawl0wski/DeployX/models"
)

type SelectScriptToEditPrompt struct {
	Project *models.Project
}

func (p *SelectScriptToEditPrompt) Run() *models.Script {
	prompt := p.preparePrompt()
	selection := p.askUser(prompt)
	return p.returnDesignAccordingToSelection(selection)

}

func (p *SelectScriptToEditPrompt) preparePrompt() survey.Prompt {
	return &survey.Select{Message: "What script do you want to edit?", Options: []string{"Before deployment script", "After deployment script"}}
}

func (p *SelectScriptToEditPrompt) askUser(prompt survey.Prompt) string {
	var selection string
	err := survey.AskOne(prompt, &selection)
	if err != nil {
		panic("Can't ask user what script to edit")
	}
	return selection
}

func (p *SelectScriptToEditPrompt) returnDesignAccordingToSelection(selection string) *models.Script {
	if selection == "Before deployment script" {
		return &p.Project.BeforeDeployScript
	} else {
		return &p.Project.AfterDeployScript
	}
}

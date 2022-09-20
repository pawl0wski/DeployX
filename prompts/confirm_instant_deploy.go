package prompts

import "github.com/AlecAivazis/survey/v2"

type ConfirmInstantDeployPrompt struct {
}

func (p *ConfirmInstantDeployPrompt) Run() bool {
	prompt := p.preparePrompt()
	return p.askUser(prompt)
}

func (p *ConfirmInstantDeployPrompt) preparePrompt() *survey.Confirm {
	return &survey.Confirm{Message: "Always deploy immediately"}
}

func (p *ConfirmInstantDeployPrompt) askUser(prompt *survey.Confirm) bool {
	var choice bool
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		panic("The instant deployment cannot be assured")
	}
	return choice
}

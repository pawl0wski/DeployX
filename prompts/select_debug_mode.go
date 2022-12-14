package prompts

import (
	"github.com/AlecAivazis/survey/v2"
)

type SelectDebugModePrompt struct {
}

func (p *SelectDebugModePrompt) Run() bool {
	prompt := p.preparePrompt()
	choice := p.askUser(prompt)
	return choice
}

func (p *SelectDebugModePrompt) preparePrompt() *survey.Confirm {
	return &survey.Confirm{Message: "Enable debug mode"}
}

func (p *SelectDebugModePrompt) askUser(prompt *survey.Confirm) bool {
	var choice bool
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		panic("Can't get debug mode")
	}
	return choice
}

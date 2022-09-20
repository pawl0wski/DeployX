package prompts

import (
	"github.com/AlecAivazis/survey/v2"
)

type GetProjectPasswordPrompt struct {
}

func (p *GetProjectPasswordPrompt) Run() string {
	prompt := p.preparePrompt()
	return p.askUser(prompt)
}

func (p *GetProjectPasswordPrompt) preparePrompt() *survey.Password {
	return &survey.Password{Message: "Password"}
}

func (p *GetProjectPasswordPrompt) askUser(prompt *survey.Password) string {
	var password string
	err := survey.AskOne(prompt, &password)
	if err != nil {
		panic("Can't get password")
	}
	return password
}

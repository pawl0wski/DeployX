package prompts

import (
	"github.com/AlecAivazis/survey/v2"
)

type GetProjectPasswordPrompt struct {
}

func (g GetProjectPasswordPrompt) Run() string {
	prompt := &survey.Password{Message: "Password"}
	var password string
	err := survey.AskOne(prompt, &password)
	if err != nil {
		panic("Can't get password")
	}
	return password
}

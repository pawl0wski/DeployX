package prompts

import (
	"github.com/AlecAivazis/survey/v2"
)

type GetProjectNamePrompt struct {
	DefaultProjectName string
}

func (p GetProjectNamePrompt) Run() string {
	prompt := &survey.Input{Message: "Name", Default: p.DefaultProjectName}
	var name string
	err := survey.AskOne(prompt, &name)
	if err != nil {
		panic("Can't get project's name")
	}
	return name
}

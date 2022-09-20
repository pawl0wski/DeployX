package prompts

import (
	"github.com/AlecAivazis/survey/v2"
)

type GetProjectNamePrompt struct {
	DefaultProjectName string
}

func (p *GetProjectNamePrompt) Run() string {
	prompt := p.preparePrompt()
	return p.askUser(prompt)
}

func (p *GetProjectNamePrompt) preparePrompt() *survey.Input {
	return &survey.Input{Message: "Name", Default: p.DefaultProjectName}
}

func (p *GetProjectNamePrompt) askUser(prompt *survey.Input) string {
	var name string
	err := survey.AskOne(prompt, &name)
	if err != nil {
		panic("Can't get project's name")
	}
	return name
}

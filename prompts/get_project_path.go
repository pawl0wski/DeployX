package prompts

import (
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"os"
)

type GetProjectPathPrompt struct {
	DefaultPath string
}

func (p *GetProjectPathPrompt) Run() string {
	prompt := p.preparePrompt()
	return p.askUser(prompt)
}

func (p *GetProjectPathPrompt) preparePrompt() *survey.Input {
	return &survey.Input{Message: "Path", Default: p.DefaultPath}
}

func (p *GetProjectPathPrompt) askUser(prompt *survey.Input) string {
	var response string
	err := survey.AskOne(prompt, &response, survey.WithValidator(p.validateProjectPath))
	if err != nil {
		panic("Can't get project's path")
	}
	return response
}

func (p *GetProjectPathPrompt) validateProjectPath(val interface{}) error {
	projectPath := p.convertValToString(val)
	_, err := os.Stat(projectPath)
	if err != nil {
		return errors.New(fmt.Sprintf("path \"%s\" not exist", val))
	}
	return nil
}

func (p *GetProjectPathPrompt) convertValToString(val interface{}) string {
	valAsString, ok := val.(string)
	if !ok {
		panic("Can't validate value")
	}
	return valAsString
}

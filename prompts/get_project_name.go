package prompts

import (
	"github.com/AlecAivazis/survey/v2"
)

func GetProjectName(defaultProjectName string) string {
	prompt := &survey.Input{Message: "Name", Default: defaultProjectName}
	var name string
	err := survey.AskOne(prompt, &name)
	if err != nil {
		panic("Can't get project's name")
	}
	return name
}

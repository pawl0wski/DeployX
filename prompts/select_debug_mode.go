package prompts

import (
	"github.com/AlecAivazis/survey/v2"
)

func SelectDebugMode() bool {
	prompt := &survey.Confirm{Message: "Enable debug mode"}
	var choice bool
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		panic("Can't get debug mode")
	}
	return choice
}
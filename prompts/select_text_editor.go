package prompts

import (
	"github.com/AlecAivazis/survey/v2"
)

func SelectTextEditor() string {
	prompt := &survey.Select{Message: "Select your favorite text editor", Options: []string{"vi", "nano", "code", "gedit", "kate"}}
	var textEditor string
	err := survey.AskOne(prompt, &textEditor)
	if err != nil {
		panic("Can't get your favorite text editor")
	}
	return textEditor
}

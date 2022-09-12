package prompts

import (
	"github.com/AlecAivazis/survey/v2"
)

func GetProjectPassword() string {
	prompt := &survey.Password{Message: "Password"}
	var password string
	err := survey.AskOne(prompt, &password)
	if err != nil {
		panic("Can't get password")
	}
	return password
}

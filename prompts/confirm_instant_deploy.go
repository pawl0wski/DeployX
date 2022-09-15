package prompts

import "github.com/AlecAivazis/survey/v2"

func ConfirmInstantDeploy() bool {
	var choice bool
	prompt := &survey.Confirm{Message: "Always deploy immediately"}
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		panic("The instant deployment cannot be assured")
	}
	return choice
}

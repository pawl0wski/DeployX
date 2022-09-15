package prompts

import "github.com/AlecAivazis/survey/v2"

func SelectDeploymentDays() []string {
	var selectedDays []string
	prompt := &survey.MultiSelect{
		Message: "Set the days on which the project can be deployed",
		Options: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"},
	}
	err := survey.AskOne(prompt, &selectedDays)
	if err != nil {
		panic("Can't ask for days")
	}
	return selectedDays
}

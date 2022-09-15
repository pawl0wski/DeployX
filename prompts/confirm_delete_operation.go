package prompts

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
)

func ConfirmDeleteOperation(whatToDelete string) bool {
	prompt := &survey.Confirm{Message: fmt.Sprintf("You sure want to delete %s?", whatToDelete), Default: false}
	var decision bool
	err := survey.AskOne(prompt, &decision)
	if err != nil {
		panic("The deletion cannot be assured")
	}
	return decision
}

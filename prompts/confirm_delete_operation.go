package prompts

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
)

type ConfirmDeleteOperationPrompt struct {
	WhatToDelete string
}

func (p ConfirmDeleteOperationPrompt) Run() bool {
	prompt := &survey.Confirm{Message: fmt.Sprintf("You sure want to delete %s?", p.WhatToDelete), Default: false}
	var decision bool
	err := survey.AskOne(prompt, &decision)
	if err != nil {
		panic("The deletion cannot be assured")
	}
	return decision
}

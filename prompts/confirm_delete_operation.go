package prompts

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
)

type ConfirmDeleteOperationPrompt struct {
	WhatToDelete string
}

func (p *ConfirmDeleteOperationPrompt) Run() bool {
	prompt := p.preparePrompt()
	return p.askUser(prompt)
}

func (p *ConfirmDeleteOperationPrompt) preparePrompt() *survey.Confirm {
	return &survey.Confirm{Message: fmt.Sprintf("You sure want to delete %s?", p.WhatToDelete), Default: false}
}

func (p *ConfirmDeleteOperationPrompt) askUser(prompt *survey.Confirm) bool {
	var decision bool
	err := survey.AskOne(prompt, &decision)
	if err != nil {
		panic("The deletion cannot be assured")
	}
	return decision
}

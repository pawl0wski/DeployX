package prompts

import (
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"strconv"
	"strings"
)

type SelectHourOptions struct {
	PromptText         string
	AllowAnyHour       bool
	GenerateHoursAfter int
}

type SelectDeploymentHourPrompt struct {
	DefaultDeployHour int
	Options           SelectHourOptions
}

func (p *SelectDeploymentHourPrompt) Run() int {
	prompt := p.preparePrompt()
	selection := p.askUser(prompt)
	hour, err := p.convertSelectionToHour(selection)
	if err != nil {
		panic(err)
	}
	return hour
}

func (p *SelectDeploymentHourPrompt) preparePrompt() *survey.Select {
	return &survey.Select{Message: p.getPromptText(), Options: p.generateHours(), Default: p.convertDeployHourToPromptDefault(p.DefaultDeployHour)}
}

func (p *SelectDeploymentHourPrompt) askUser(prompt *survey.Select) string {
	var selection string
	err := survey.AskOne(prompt, &selection)
	if err != nil {
		panic("Can't select deployment hour")
	}
	return selection
}

func (p *SelectDeploymentHourPrompt) convertSelectionToHour(selection string) (int, error) {
	if selection == "Any hour" {
		return -1, nil
	}
	var splitHour = strings.Split(selection, ":")
	hour, err := strconv.Atoi(splitHour[0])
	if err != nil {
		return hour, errors.New(fmt.Sprintf("Can't transform \"%s\" to int", splitHour[0]))
	}
	return hour, nil
}

func (p *SelectDeploymentHourPrompt) convertDeployHourToPromptDefault(deployHour int) string {
	if deployHour == -1 {
		return "Any hour"
	}
	return fmt.Sprintf("%d:00", deployHour)
}

func (p *SelectDeploymentHourPrompt) generateHours() []string {
	var hours []string
	if p.Options.AllowAnyHour {
		hours = append(hours, "Any hour")
	}
	for i := p.Options.GenerateHoursAfter; i < 24; i++ {
		hours = append(hours, fmt.Sprintf("%d:00", i))
	}
	return hours
}

func (p *SelectDeploymentHourPrompt) getPromptText() string {
	text := p.Options.PromptText
	if text == "" {
		return "Deploy at"
	}
	return text
}

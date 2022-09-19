package prompts

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"strconv"
	"strings"
)

type SelectDeploymentHourPrompt struct {
	DefaultDeployHour int
}

func (p SelectDeploymentHourPrompt) Run() int {
	var stringHour string
	prompt := &survey.Select{Message: "Deploy at", Options: p.generateHours(), Default: p.convertDeployHourToPromptDefault(p.DefaultDeployHour)}
	err := survey.AskOne(prompt, &stringHour)
	if err != nil {
		panic("Can't select deployment hour")
	}
	if stringHour == "Any hour" {
		return -1
	}
	var splitHour = strings.Split(stringHour, ":")
	var hour int
	hour, err = strconv.Atoi(splitHour[0])
	if err != nil {
		panic(fmt.Sprintf("Can't transform \"%s\" to int", splitHour[0]))
	}
	return hour
}

func (p SelectDeploymentHourPrompt) convertDeployHourToPromptDefault(deployHour int) string {
	if deployHour == -1 {
		return "Any hour"
	}
	return fmt.Sprintf("%d:00", deployHour)
}

func (p SelectDeploymentHourPrompt) generateHours() []string {
	var hours []string
	hours = append(hours, "Any hour")
	for i := 0; i < 24; i++ {
		hours = append(hours, fmt.Sprintf("%d:00", i))
	}
	return hours
}

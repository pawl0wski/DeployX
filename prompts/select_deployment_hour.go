package prompts

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"strconv"
	"strings"
)

func SelectDeploymentHour(defaultDeployHour int) int {
	var stringHour string
	prompt := &survey.Select{Message: "Deploy at", Options: generateHours(), Default: convertDeployHourToPromptDefault(defaultDeployHour)}
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

func convertDeployHourToPromptDefault(deployHour int) string {
	if deployHour == -1 {
		return "Any hour"
	}
	return fmt.Sprintf("%d:00", deployHour)
}

func generateHours() []string {
	var hours []string
	hours = append(hours, "Any hour")
	for i := 0; i < 24; i++ {
		hours = append(hours, fmt.Sprintf("%d:00", i))
	}
	return hours
}

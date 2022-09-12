package prompts

import (
	"DeployX/models"
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"strconv"
)

func GetServerPort(config *models.Config) uint16 {
	if config.ServerPort == 0 {
		config.ServerPort = 7777
	}
	prompt := &survey.Input{Message: "Server port", Default: strconv.Itoa(int(config.ServerPort))}
	var serverPortAsString string
	err := survey.AskOne(prompt, &serverPortAsString, survey.WithValidator(validateServerPort))
	if err != nil {
		panic("Can't get server port")
	}
	port, err := strconv.ParseUint(serverPortAsString, 10, 16)
	if err != nil {
		panic("Invalid port")
	}
	return uint16(port)
}

func validateServerPort(val interface{}) error {
	serverPort := convertValToString(val)
	_, err := strconv.ParseUint(serverPort, 10, 16)
	if err != nil {
		return errors.New(fmt.Sprintf("%s is not valid port", serverPort))
	}
	return nil
}

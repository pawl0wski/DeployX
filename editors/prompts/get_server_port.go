package prompts

import (
	"DeployX/models"
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"strconv"
)

func GetServerPort(config *models.Config) uint16 {
	if config.ServerPort == 0 {
		config.ServerPort = 7777
	}
	prompt := promptui.Prompt{Label: "Server port", Default: strconv.Itoa(int(config.ServerPort)), Validate: validateServerPort}
	result, err := prompt.Run()
	if err != nil {
		panic("Can't get server port")
	}
	port, err := strconv.ParseUint(result, 10, 16)
	if err != nil {
		panic("Invalid port")
	}
	return uint16(port)
}

func validateServerPort(input string) error {
	_, err := strconv.ParseUint(input, 10, 16)
	if err != nil {
		return errors.New(fmt.Sprintf("%s is not valid port", input))
	}
	return nil
}

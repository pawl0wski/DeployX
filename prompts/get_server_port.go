package prompts

import (
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"strconv"
)

type GetServerPortPrompt struct {
	DefaultPort uint16
}

func (p *GetServerPortPrompt) Run() uint16 {
	defaultPort := p.getDefaultPort()
	prompt := &survey.Input{Message: "Server port", Default: strconv.Itoa(int(defaultPort))}
	response := p.askUser(prompt)
	port := p.convertResponseToUInt16(response)
	return port
}

func (p *GetServerPortPrompt) getDefaultPort() uint16 {
	if p.DefaultPort == 0 {
		return 7777
	}
	return p.DefaultPort
}

func (p *GetServerPortPrompt) preparePrompt(defaultPort uint16) *survey.Input {
	return &survey.Input{Message: "Server port", Default: strconv.Itoa(int(defaultPort))}
}

func (p *GetServerPortPrompt) askUser(prompt *survey.Input) string {
	var response string
	err := survey.AskOne(prompt, &response, survey.WithValidator(p.validateServerPort))
	if err != nil {
		panic("Can't get server port")
	}
	return response
}

func (p *GetServerPortPrompt) convertResponseToUInt16(response string) uint16 {
	port, err := strconv.ParseUint(response, 10, 16)
	if err != nil {
		panic("Invalid port")
	}
	return uint16(port)
}

func (p *GetServerPortPrompt) validateServerPort(val interface{}) error {
	serverPort := p.convertValToString(val)
	_, err := strconv.ParseUint(serverPort, 10, 16)
	if err != nil {
		return errors.New(fmt.Sprintf("%s is not valid port", serverPort))
	}
	return nil
}

func (p *GetServerPortPrompt) convertValToString(val interface{}) string {
	valAsString, ok := val.(string)
	if !ok {
		panic("Can't validate value")
	}
	return valAsString
}

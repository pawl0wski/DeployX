package models

import (
	"fmt"
	"os/exec"
	"strings"
)

type Command struct {
	Command          string
	DefaultArguments []string
}

func (c Command) CheckIfExist() bool {
	_, err := exec.LookPath(c.Command)
	if err != nil {
		return false
	}
	return true
}

func (c Command) ToString() string {
	return fmt.Sprintf("%s %s", c.Command, strings.Join(c.DefaultArguments, " "))
}

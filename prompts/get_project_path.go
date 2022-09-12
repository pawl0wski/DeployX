package prompts

import (
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"os"
)

func GetProjectPath(defaultPath string) string {
	prompt := &survey.Input{Message: "Path", Default: defaultPath}
	var projectPath string
	err := survey.AskOne(prompt, &projectPath, survey.WithValidator(validateProjectPath))
	if err != nil {
		panic("Can't get project's path")
	}
	return projectPath
}

func validateProjectPath(val interface{}) error {
	projectPath := convertValToString(val)
	_, err := os.Stat(projectPath)
	if err != nil {
		return errors.New(fmt.Sprintf("path \"%s\" not exist", val))
	}
	return nil
}

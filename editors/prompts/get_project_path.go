package prompts

import (
	"DeployX/models"
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"os"
)

func GetProjectPath(project *models.Project) string {
	prompt := &survey.Input{Message: "Path", Default: project.Path}
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

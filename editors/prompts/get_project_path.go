package prompts

import (
	"DeployX/models"
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
)

func GetProjectPath(project *models.Project) string {
	prompt := promptui.Prompt{Label: "Path", Validate: validateProjectPath, Default: project.Path}
	result, err := prompt.Run()
	if err != nil {
		panic("Can't get project's path")
	}
	return result
}

func validateProjectPath(input string) error {
	_, err := os.Stat(input)
	if err != nil {
		return errors.New(fmt.Sprintf("path \"%s\" not exist", input))
	}
	return nil
}

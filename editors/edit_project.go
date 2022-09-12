package editors

import (
	"DeployX/models"
	"DeployX/prompts"
)

func EditProject(project *models.Project) {
	project.Name = prompts.GetProjectName(project.Name)
	project.Path = prompts.GetProjectPath(project.Path)
	project.SetPassword(prompts.GetProjectPassword())
}

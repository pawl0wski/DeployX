package editors

import (
	"DeployX/editors/prompts"
	"DeployX/models"
)

func EditProject(project *models.Project) {
	printWithEditorColor("Project editor")
	project.Name = prompts.GetProjectName(project)
	project.Path = prompts.GetProjectPath(project)
	project.SetPassword(prompts.GetProjectPassword())
}

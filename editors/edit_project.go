package editors

import (
	"DeployX/models"
	prompts2 "DeployX/prompts"
)

func EditProject(project *models.Project) {
	printWithEditorColor("Project editor")
	project.Name = prompts2.GetProjectName(project)
	project.Path = prompts2.GetProjectPath(project)
	project.SetPassword(prompts2.GetProjectPassword())
}

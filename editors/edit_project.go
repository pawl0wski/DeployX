package editors

import (
	"github.com/pawl0wski/DeployX/models"
	"github.com/pawl0wski/DeployX/prompts"
)

func EditProject(project *models.Project) {
	projectNamePrompt := prompts.GetProjectNamePrompt{DefaultProjectName: project.Name}
	project.Name = projectNamePrompt.Run()
	projectPathPrompt := prompts.GetProjectPathPrompt{DefaultPath: project.Path}
	project.Path = projectPathPrompt.Run()
	projectPasswordPrompt := prompts.GetProjectPasswordPrompt{}
	project.SetPassword(projectPasswordPrompt.Run())
	deploymentConfiguration := project.DeploymentConfiguration
	EditDeploymentConfiguration(&deploymentConfiguration)
	project.DeploymentConfiguration = deploymentConfiguration
}

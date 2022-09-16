package editors

import (
	"github.com/pawl0wski/DeployX/models"
	"github.com/pawl0wski/DeployX/prompts"
)

func EditProject(project *models.Project) {
	project.Name = prompts.GetProjectName(project.Name)
	project.Path = prompts.GetProjectPath(project.Path)
	project.SetPassword(prompts.GetProjectPassword())
	deploymentConfiguration := project.DeploymentConfiguration
	EditDeploymentConfiguration(&deploymentConfiguration)
	project.DeploymentConfiguration = deploymentConfiguration
}

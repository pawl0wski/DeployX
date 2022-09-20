package editors

import (
	"github.com/pawl0wski/DeployX/models"
	"github.com/pawl0wski/DeployX/prompts"
)

type ProjectEditor struct {
	ProjectToEdit *models.Project
}

func (e *ProjectEditor) Run() {
	e.editProjectName()
	e.editProjectPath()
	e.editProjectPassword()
	e.editDeploymentConfiguration()
}

func (e *ProjectEditor) editProjectName() {
	projectNamePrompt := prompts.GetProjectNamePrompt{DefaultProjectName: e.ProjectToEdit.Name}
	e.ProjectToEdit.Name = projectNamePrompt.Run()
}

func (e *ProjectEditor) editProjectPath() {
	projectPathPrompt := prompts.GetProjectPathPrompt{DefaultPath: e.ProjectToEdit.Path}
	e.ProjectToEdit.Path = projectPathPrompt.Run()
}

func (e *ProjectEditor) editProjectPassword() {
	projectPasswordPrompt := prompts.GetProjectPasswordPrompt{}
	e.ProjectToEdit.SetPassword(projectPasswordPrompt.Run())
}

func (e *ProjectEditor) editDeploymentConfiguration() {
	deploymentConfiguration := e.ProjectToEdit.DeploymentConfiguration
	deploymentConfigurationEditor := DeploymentConfigurationEditor{ConfigurationToEdit: &deploymentConfiguration}
	deploymentConfigurationEditor.Run()
	e.ProjectToEdit.DeploymentConfiguration = deploymentConfiguration
}

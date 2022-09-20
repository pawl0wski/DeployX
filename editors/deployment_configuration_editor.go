package editors

import (
	"github.com/pawl0wski/DeployX/models"
	"github.com/pawl0wski/DeployX/prompts"
)

type DeploymentConfigurationEditor struct {
	ConfigurationToEdit *models.DeploymentConfiguration
}

func (e *DeploymentConfigurationEditor) Run() {
	e.editInstantDeploy()
	if !e.ConfigurationToEdit.Instant {
		e.editDeployAfterHour()
		if e.ConfigurationToEdit.DeployAfterHour != -1 {
			e.editDeployBeforeHour()
		}
		e.editDeployDays()
	}
}

func (e *DeploymentConfigurationEditor) editInstantDeploy() {
	instantDeployPrompt := prompts.ConfirmInstantDeployPrompt{}
	e.ConfigurationToEdit.Instant = instantDeployPrompt.Run()
}

func (e *DeploymentConfigurationEditor) editDeployAfterHour() {
	deployAfterHourPrompt := prompts.SelectDeploymentHourPrompt{DefaultDeployHour: e.ConfigurationToEdit.DeployAfterHour, Options: prompts.SelectHourOptions{
		AllowAnyHour:       true,
		GenerateHoursAfter: 0,
		PromptText:         "Deploy after",
	}}
	e.ConfigurationToEdit.DeployAfterHour = deployAfterHourPrompt.Run()
}

func (e *DeploymentConfigurationEditor) editDeployBeforeHour() {
	deployBeforeHourPrompt := prompts.SelectDeploymentHourPrompt{DefaultDeployHour: e.ConfigurationToEdit.DeployBeforeHour, Options: prompts.SelectHourOptions{AllowAnyHour: false, GenerateHoursAfter: e.ConfigurationToEdit.DeployAfterHour, PromptText: "Deploy before"}}
	e.ConfigurationToEdit.DeployBeforeHour = deployBeforeHourPrompt.Run()
}

func (e *DeploymentConfigurationEditor) editDeployDays() {
	weekdaysPrompt := prompts.SelectWeekdaysPrompt{DefaultWeekdays: e.ConfigurationToEdit.GetWeekdays()}
	e.ConfigurationToEdit.SaveWeekdays(weekdaysPrompt.Run())
}

package editors

import (
	"github.com/pawl0wski/DeployX/models"
	"github.com/pawl0wski/DeployX/prompts"
	"time"
)

func EditDeploymentConfiguration(configuration *models.DeploymentConfiguration) {
	instantDeployPrompt := prompts.ConfirmInstantDeployPrompt{}
	instantDeploy := instantDeployPrompt.Run()
	setInstantDeploy(instantDeploy, configuration)
	if !instantDeploy {
		deployBeforeHour := 0
		deployAfterHourPrompt := prompts.SelectDeploymentHourPrompt{DefaultDeployHour: configuration.DeployAfterHour, Options: prompts.SelectHourOptions{
			AllowAnyHour:       true,
			GenerateHoursAfter: 0,
			PromptText:         "Deploy after",
		}}
		deployAfterHour := deployAfterHourPrompt.Run()
		if deployAfterHour != -1 {
			deployBeforeHourPrompt := prompts.SelectDeploymentHourPrompt{DefaultDeployHour: configuration.DeployBeforeHour, Options: prompts.SelectHourOptions{AllowAnyHour: false, GenerateHoursAfter: deployAfterHour, PromptText: "Deploy before"}}
			deployBeforeHour = deployBeforeHourPrompt.Run()
		}
		weekdaysPrompt := prompts.SelectWeekdaysPrompt{DefaultWeekdays: configuration.GetWeekdays()}
		weekdays := weekdaysPrompt.Run()
		setDeployAfterHour(deployAfterHour, configuration)
		setDeployBeforeHour(deployBeforeHour, configuration)
		setDeployDays(weekdays, configuration)
	}
}

func setInstantDeploy(instantDeploy bool, configuration *models.DeploymentConfiguration) {
	configuration.Instant = instantDeploy
}

func setDeployBeforeHour(deployBeforeHour int, configuration *models.DeploymentConfiguration) {
	configuration.DeployBeforeHour = deployBeforeHour
}

func setDeployAfterHour(deployAfterHour int, configuration *models.DeploymentConfiguration) {
	configuration.DeployAfterHour = deployAfterHour
}

func setDeployDays(weekdays []time.Weekday, configuration *models.DeploymentConfiguration) {
	configuration.SaveWeekdays(weekdays)
}

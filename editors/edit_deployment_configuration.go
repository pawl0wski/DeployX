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
		deployHourPrompt := prompts.SelectDeploymentHourPrompt{DefaultDeployHour: configuration.DeployAfterHour}
		deployHour := deployHourPrompt.Run()
		weekdaysPrompt := prompts.SelectWeekdaysPrompt{DefaultWeekdays: configuration.DeserializeAndReturnWeekdays()}
		weekdays := weekdaysPrompt.Run()
		setDeployHour(deployHour, configuration)
		setDeployDays(weekdays, configuration)
	}
}

func setInstantDeploy(instantDeploy bool, configuration *models.DeploymentConfiguration) {
	configuration.Instant = instantDeploy
}

func setDeployHour(deployHour int, configuration *models.DeploymentConfiguration) {
	configuration.DeployAfterHour = deployHour
}

func setDeployDays(weekdays []time.Weekday, configuration *models.DeploymentConfiguration) {
	configuration.SerializeAndSaveWeekdays(weekdays)
}

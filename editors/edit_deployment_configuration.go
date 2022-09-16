package editors

import (
	"github.com/pawl0wski/DeployX/models"
	"github.com/pawl0wski/DeployX/prompts"
	"time"
)

func EditDeploymentConfiguration(configuration *models.DeploymentConfiguration) {
	instantDeploy := prompts.ConfirmInstantDeploy()
	setInstantDeploy(instantDeploy, configuration)
	if !instantDeploy {
		deployHour := prompts.SelectDeploymentHour(configuration.DeployAfterHour)
		weekdays := prompts.SelectWeekdays(configuration.DeserializeAndReturnWeekdays())
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

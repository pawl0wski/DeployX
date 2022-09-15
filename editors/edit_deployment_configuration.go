package editors

import (
	"github.com/pawl0wski/DeployX/models"
	"github.com/pawl0wski/DeployX/prompts"
)

func EditDeploymentConfiguration(configuration *models.DeploymentConfiguration) {
	instantDeploy := prompts.ConfirmInstantDeploy()
	setInstantDeploy(instantDeploy, configuration)
	if !instantDeploy {
		deployHour := prompts.SelectDeploymentHour()
		days := prompts.SelectDeploymentDays()
		setDeployHour(deployHour, configuration)
		setDeployDays(days, configuration)
	}
}

func setInstantDeploy(instantDeploy bool, configuration *models.DeploymentConfiguration) {
	configuration.Instant = instantDeploy
}

func setDeployHour(deployHour int, configuration *models.DeploymentConfiguration) {
	configuration.DeployAfterHour = deployHour
}

func setDeployDays(days []string, configuration *models.DeploymentConfiguration) {
	for _, day := range days {
		switch day {
		case "Monday":
			configuration.DeployOnMonday = true
			break
		case "Tuesday":
			configuration.DeployOnTuesday = true
			break
		case "Wednesday":
			configuration.DeployOnWednesday = true
			break
		case "Thursday":
			configuration.DeployOnThursday = true
			break
		case "Friday":
			configuration.DeployOnFriday = true
			break
		case "Saturday":
			configuration.DeployOnSaturday = true
			break
		case "Sunday":
			configuration.DeployOnSunday = true
			break
		}
	}
}

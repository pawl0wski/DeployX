package models

import (
	"github.com/pawl0wski/DeployX/database"
	"gorm.io/gorm"
)

type DeploymentConfiguration struct {
	gorm.Model
	Instant           bool `json:"instant"`
	DeployAfterHour   int  `json:"deploy_after_hour"`
	DeployOnMonday    bool `json:"deploy_on_monday"`
	DeployOnTuesday   bool `json:"deploy_on_tuesday"`
	DeployOnWednesday bool `json:"deploy_on_wednesday"`
	DeployOnThursday  bool `json:"deploy_on_thursday"`
	DeployOnFriday    bool `json:"deploy_on_friday"`
	DeployOnSaturday  bool `json:"deploy_on_saturday"`
	DeployOnSunday    bool `json:"deploy_on_sunday"`
}

func (configuration *DeploymentConfiguration) SetFalseForAllDays() {
	configuration.DeployOnMonday = false
	configuration.DeployOnTuesday = false
	configuration.DeployOnWednesday = false
	configuration.DeployOnThursday = false
	configuration.DeployOnFriday = false
	configuration.DeployOnSaturday = false
}

func InitializeDeploymentConfiguration() {
	database.InitializeModel(DeploymentConfiguration{})
}

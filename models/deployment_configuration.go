package models

type DeploymentConfiguration struct {
	Instant           bool `json:"instant"`
	DeployAfter       byte `json:"deploy_after"`
	DeployOnMonday    bool `json:"deploy_on_monday"`
	DeployOnTuesday   bool `json:"deploy_on_tuesday"`
	DeployOnWednesday bool `json:"deploy_on_wednesday"`
	DeployOnThursday  bool `json:"deploy_on_thursday"`
	DeployOnFriday    bool `json:"deploy_on_friday"`
	DeployOnSaturday  bool `json:"deploy_on_saturday"`
	DeployOnSunday    bool `json:"deploy_on_sunday"`
}

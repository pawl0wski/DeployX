package models

import (
	"fmt"
	"github.com/pawl0wski/DeployX/database"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type DeploymentConfiguration struct {
	gorm.Model
	Instant          bool   `json:"instant"`
	DeployAfterHour  int    `json:"deploy_after_hour"`
	DeployOnWeekdays string `json:"deploy_on_weekdays"`
}

func (configuration *DeploymentConfiguration) SerializeAndSaveWeekdays(weekdays []time.Weekday) {
	var serializedWeekdays string
	for _, weekday := range weekdays {
		serializedWeekdays += fmt.Sprintf("%d,", int(weekday))
	}
	configuration.DeployOnWeekdays = strings.TrimRight(serializedWeekdays, ",")
}

func (configuration *DeploymentConfiguration) DeserializeAndReturnWeekdays() []time.Weekday {
	var weekdays []time.Weekday
	if configuration.DeployOnWeekdays == "" {
		return weekdays
	}
	for _, weekday := range strings.Split(configuration.DeployOnWeekdays, ",") {
		intWeekday, err := strconv.Atoi(weekday)
		if err != nil {
			panic(err)
		}
		weekdays = append(weekdays, time.Weekday(intWeekday))
	}
	return weekdays
}

func (configuration *DeploymentConfiguration) ClearWeekDays() {
	configuration.DeployOnWeekdays = ""
}

func InitializeDeploymentConfiguration() {
	database.InitializeModel(DeploymentConfiguration{})
}

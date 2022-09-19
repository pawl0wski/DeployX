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
	Instant            bool   `json:"instant"`
	DeployAfterHour    int    `json:"deploy_after_hour"`
	DeployBeforeHour   int    `json:"deploy_before_hour"`
	SerializedWeekdays string `json:"serialized_weekdays"`
}

func (configuration *DeploymentConfiguration) SaveWeekdays(weekdays []time.Weekday) {
	var serializedWeekdays string
	for _, weekday := range weekdays {
		serializedWeekdays += fmt.Sprintf("%d,", int(weekday))
	}
	configuration.SerializedWeekdays = strings.TrimRight(serializedWeekdays, ",")
}

func (configuration *DeploymentConfiguration) GetWeekdays() []time.Weekday {
	var weekdays []time.Weekday
	if configuration.SerializedWeekdays == "" {
		return weekdays
	}
	for _, weekday := range strings.Split(configuration.SerializedWeekdays, ",") {
		intWeekday, err := strconv.Atoi(weekday)
		if err != nil {
			panic(err)
		}
		weekdays = append(weekdays, time.Weekday(intWeekday))
	}
	return weekdays
}

func (configuration *DeploymentConfiguration) ClearWeekDays() {
	configuration.SerializedWeekdays = ""
}

func InitializeDeploymentConfiguration() {
	database.InitializeModel(DeploymentConfiguration{})
}

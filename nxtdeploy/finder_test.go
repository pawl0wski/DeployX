package nxtdeploy_test

import (
	"github.com/pawl0wski/DeployX/models"
	"github.com/pawl0wski/DeployX/nxtdeploy"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFinderFindDeployTimeToday(t *testing.T) {
	deploymentConf := models.DeploymentConfiguration{
		DeployAfterHour:    10,
		DeployBeforeHour:   20,
		SerializedWeekdays: "0",
	}
	today := time.Date(2020, 1, 12, 10, 53, 23, 14, time.UTC)
	expectedDeployTime := time.Date(2020, 1, 12, 11, 0, 0, 0, time.UTC)
	f := nxtdeploy.Finder{DeploymentConfiguration: deploymentConf}

	deploymentDate, err := f.FindDeployTime(today)

	assert.NoError(t, err)
	assert.Equal(t, &expectedDeployTime, deploymentDate)
}

func TestFinderFindDeployTimeTomorrow(t *testing.T) {
	deploymentConf := models.DeploymentConfiguration{
		DeployAfterHour:    0,
		DeployBeforeHour:   0,
		SerializedWeekdays: "0",
	}
	today := time.Date(2022, 11, 12, 15, 32, 23, 2, time.UTC)
	expectedDeployTime := time.Date(2022, 11, 13, 0, 0, 0, 0, time.UTC)
	f := nxtdeploy.Finder{DeploymentConfiguration: deploymentConf}

	deploymentDate, err := f.FindDeployTime(today)

	assert.NoError(t, err)
	assert.Equal(t, &expectedDeployTime, deploymentDate)
}

func TestFinderThrowErrorIfWeekPassed(t *testing.T) {
	deploymentConf := models.DeploymentConfiguration{
		DeployAfterHour:    0,
		DeployBeforeHour:   0,
		SerializedWeekdays: "",
	}
	today := time.Date(2022, 11, 12, 10, 10, 10, 10, time.UTC)
	f := nxtdeploy.Finder{DeploymentConfiguration: deploymentConf}

	_, err := f.FindDeployTime(today)

	assert.Error(t, err)
}

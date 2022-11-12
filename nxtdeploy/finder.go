package nxtdeploy

import (
	"errors"
	"github.com/pawl0wski/DeployX/models"
	"time"
)

type Finder struct {
	deployTime              *time.Time
	DeploymentConfiguration models.DeploymentConfiguration
}

func (f *Finder) FindDeployTime(fromTime time.Time) (*time.Time, error) {
	f.deployTime = &fromTime
	f.resetMinAndSec()
	hoursPassed := 0
	const HoursInWeek = 24 * 7
	for !f.checkIfDeployTimeIsCorrect() {
		if hoursPassed >= HoursInWeek {
			return f.deployTime, errors.New("can't find any deployment date")
		}
		f.shiftDeployTimeByHour()
		hoursPassed++
	}
	return f.deployTime, nil
}

func (f *Finder) resetMinAndSec() {
	t := f.deployTime
	cleanedTime := time.Date(t.Year(), t.Month(), t.Day(), t.Hour()+1, 0, 0, 0, t.Location())
	f.deployTime = &cleanedTime
}

func (f *Finder) checkIfDeployTimeIsCorrect() bool {
	t := f.deployTime
	correctDeployAfterHour := t.Hour() >= f.DeploymentConfiguration.DeployAfterHour
	correctDeployBeforeHour := t.Hour() <= f.DeploymentConfiguration.DeployBeforeHour
	deployWeekdays := f.DeploymentConfiguration.GetWeekdays()
	correctWeekday := false
	for _, weekday := range deployWeekdays {
		if weekday.String() == t.Weekday().String() {
			correctWeekday = true
		}
	}
	return correctDeployAfterHour && correctDeployBeforeHour && correctWeekday
}

func (f *Finder) shiftDeployTimeByHour() {
	shiftedTime := f.deployTime.Add(time.Hour)
	f.deployTime = &shiftedTime
}

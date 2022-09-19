package prompts

import (
	"errors"
	"github.com/AlecAivazis/survey/v2"
	"time"
)

type SelectWeekdaysPrompt struct {
	DefaultWeekdays []time.Weekday
}

func (p SelectWeekdaysPrompt) Run() []time.Weekday {
	var selections []string
	prompt := &survey.MultiSelect{
		Default: p.convertWeekdaysToStrings(p.DefaultWeekdays),
		Message: "Set the days on which the project can be deployed",
		Options: p.convertWeekdaysToStrings(p.generateWeekdays()),
	}
	err := survey.AskOne(prompt, &selections)
	if err != nil {
		panic("Can't ask for days")
	}
	var selectedWeekdays []time.Weekday
	selectedWeekdays, err = p.convertWeekdayStringsToWeekdays(selections)
	if err != nil {
		panic(err)
	}
	return selectedWeekdays
}

func (p SelectWeekdaysPrompt) generateWeekdays() []time.Weekday {
	var weekdays []time.Weekday
	for i := 0; i < 7; i++ {
		weekdays = append(weekdays, time.Weekday(i))
	}
	return weekdays
}

func (p SelectWeekdaysPrompt) convertWeekdaysToStrings(weekdays []time.Weekday) []string {
	var weekdayStrings []string
	for _, weekday := range weekdays {
		weekdayStrings = append(weekdayStrings, weekday.String())
	}
	return weekdayStrings
}

func (p SelectWeekdaysPrompt) convertWeekdayStringsToWeekdays(weekdayStrings []string) ([]time.Weekday, error) {
	var weekdays []time.Weekday
	for _, weekdayString := range weekdayStrings {
		weekdayIndex, err := p.convertWeekdayStringToIndex(weekdayString)
		if err != nil {
			return weekdays, err
		}
		weekdays = append(weekdays, time.Weekday(weekdayIndex))
	}
	return weekdays, nil
}

func (p SelectWeekdaysPrompt) convertWeekdayStringToIndex(weekdayString string) (int, error) {
	for i, weekday := range p.convertWeekdaysToStrings(p.generateWeekdays()) {
		if weekday == weekdayString {
			return i, nil
		}
	}
	return 0, errors.New("invalid weekday name")
}

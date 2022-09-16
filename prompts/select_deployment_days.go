package prompts

import (
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"time"
)

func SelectWeekdays(defaultWeekdays []time.Weekday) []time.Weekday {
	var selections []string
	fmt.Println(convertWeekdaysToStrings(defaultWeekdays))
	prompt := &survey.MultiSelect{
		Default: convertWeekdaysToStrings(defaultWeekdays),
		Message: "Set the days on which the project can be deployed",
		Options: convertWeekdaysToStrings(generateWeekdays()),
	}
	err := survey.AskOne(prompt, &selections)
	if err != nil {
		panic("Can't ask for days")
	}
	var selectedWeekdays []time.Weekday
	selectedWeekdays, err = convertWeekdayStringsToWeekdays(selections)
	if err != nil {
		panic(err)
	}
	return selectedWeekdays
}

func generateWeekdays() []time.Weekday {
	var weekdays []time.Weekday
	for i := 0; i < 7; i++ {
		weekdays = append(weekdays, time.Weekday(i))
	}
	return weekdays
}

func convertWeekdaysToStrings(weekdays []time.Weekday) []string {
	var weekdayStrings []string
	for _, weekday := range weekdays {
		weekdayStrings = append(weekdayStrings, weekday.String())
	}
	return weekdayStrings
}

func convertWeekdayStringsToWeekdays(weekdayStrings []string) ([]time.Weekday, error) {
	var weekdays []time.Weekday
	for _, weekdayString := range weekdayStrings {
		weekdayIndex, err := convertWeekdayStringToIndex(weekdayString)
		if err != nil {
			return weekdays, err
		}
		weekdays = append(weekdays, time.Weekday(weekdayIndex))
	}
	return weekdays, nil
}

func convertWeekdayStringToIndex(weekdayString string) (int, error) {
	for i, weekday := range convertWeekdaysToStrings(generateWeekdays()) {
		if weekday == weekdayString {
			return i, nil
		}
	}
	return 0, errors.New("invalid weekday name")
}

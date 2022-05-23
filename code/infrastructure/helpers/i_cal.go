package helpers

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"strings"
)

var daysMap = map[string]float64{
	"SU": 0,
	"MO": 1,
	"TU": 2,
	"WE": 3,
	"TH": 4,
	"FR": 5,
	"SA": 6,
}

func GetRecurrenceRuleData(recurrenceRule, startDateString string) (weekdays []interface{}, until string, err error) {
	recurrenceRuleParts := strings.Split(recurrenceRule, ";")
	var recurrenceRules = make(map[string]string)
	for _, part := range recurrenceRuleParts {
		rule := strings.Split(part, "=")
		recurrenceRules[rule[0]] = rule[1]
	}
	until = recurrenceRules["UNTIL"]
	err = validation.Validate(until, Date("20060102T150405Z"), StringDateAfterString(startDateString, "2006-01-02T15:04:05.000Z", "20060102T150405Z", "startDate"))
	if err != nil {
		return nil, "", err
	}
	days := strings.Split(recurrenceRules["BYDAY"], ",")
	for _, day := range days {
		err = validation.Validate(day, In("SU", "MO", "TU", "WE", "TH", "FR", "SA"))
		if err != nil {
			return nil, "", err
		}
		weekdays = append(weekdays, daysMap[day])
	}
	return weekdays, until, err
}

package helpers

import (
	"errors"
	"fmt"
	"github.com/bykovme/gotrans"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"strings"
	"time"
)

func When(conditions bool, rule ...validation.Rule) validation.Rule {
	return validation.When(conditions, rule...)
}

func Required() validation.Rule {
	return validation.Required.Error(gotrans.T("rules.required"))
}

func Date(layout string) validation.Rule {
	return validation.Date(layout).Error(fmt.Sprintf(gotrans.T("rules.date"), layout))
}

func In(values ...interface{}) validation.Rule {
	return validation.In(values...).Error(fmt.Sprintf(gotrans.T("rules.in"), values))
}

func DateAfter(min time.Time, dateName string) validation.Rule {
	return validation.Min(min).Error(fmt.Sprintf(gotrans.T("rules.date_after"), dateName))
}

func DateBefore(min time.Time, dateName string) validation.Rule {
	return validation.Max(min).Error(fmt.Sprintf(gotrans.T("rules.date_before"), dateName))
}

func UUID() validation.Rule {
	return is.UUID.Error(gotrans.T("rules.uuid"))
}

func IsICalRecurrenceRule() validation.Rule {
	return validation.By(func(value interface{}) error {
		svalue := value.(string)
		if !(strings.Contains(svalue, "FREQ=WEEKLY") && strings.Contains(svalue, "UNTIL=") && strings.Contains(svalue, "BYDAY=")) {
			return fmt.Errorf(gotrans.T("rules.i_cal_recurrence_rule"))
		}
		return nil
	})
}

func IsFloat64Slice() validation.Rule {
	return validation.By(func(value interface{}) error {
		values := value.([]interface{})
		for _, item := range values {
			_, isFloat64Slice := item.(float64)
			if !isFloat64Slice {
				return errors.New(gotrans.T("float_slice"))
			}
		}
		return nil
	})
}

func StringDateAfter(min time.Time, layout, dateName string) validation.Rule {
	return validation.By(func(value interface{}) error {
		stringDate := value.(string)
		date, err := time.Parse(layout, stringDate)
		if err != nil {
			return err
		}
		if !(date.After(min) || date.Equal(min)) {
			return fmt.Errorf(gotrans.T("rules.date_after"), dateName)
		}
		return nil
	})
}

func StringDateAfterString(min, minLayout, layout, dateName string) validation.Rule {
	return validation.By(func(value interface{}) error {
		stringDate := value.(string)
		date, err := time.Parse(layout, stringDate)
		minDate, err := time.Parse(minLayout, min)
		if err != nil {
			return err
		}
		if !(date.After(minDate) || date.Equal(minDate)) {
			return fmt.Errorf(gotrans.T("rules.date_after"), dateName)
		}
		return nil
	})
}

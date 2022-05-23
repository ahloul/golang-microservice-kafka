package helpers

import (
	"fmt"
	"github.com/bykovme/gotrans"
	"github.com/ahloul/loyalty-reports/infrastructure/config"
	"strings"
)

type App string

const (
	Timetable App = "/timetable/api/v1/:lang"
	Dashboard App = "/dashboard/api/v1/:lang"
)

type Endpoint string

const (
	DashboardSubjectsLookup  Endpoint = "/employee/classrooms/:classroom_uuid/subjects/lookup"
	DashboardTeachersLookup  Endpoint = "/employee/classrooms/:classroom_uuid/available-teachers/lookup"
	TimetableCreateTimetable Endpoint = "/employee/classrooms/:classroom_uuid/timetable"
	SessionData              Endpoint = "/employee/classrooms/:classroom_uuid/timetable/:timetable_uuid/sessions"
)

func BuildEndpointUrl(app App, endpoint Endpoint, params ...map[string]interface{}) string {
	lang := gotrans.GetDefaultLocale()
	viper := config.NewViper()
	url := fmt.Sprint(viper.Server.BaseHost, app, endpoint)
	url = strings.ReplaceAll(url, ":lang", lang)
	if len(params) != 0 {
		for key, value := range params[0] {
			if strings.Contains(url, ":"+key) {
				url = strings.ReplaceAll(url, ":"+key, fmt.Sprint(value))
			} else if strings.Contains(url, "?") {
				url += fmt.Sprint("&", key, "=", value)
			} else {
				url += fmt.Sprint("?", key, "=", value)
			}
		}
	}
	return url
}

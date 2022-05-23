package ports

type PurchaseResponse struct {
	Id                  string      `jsonapi:"primary,timetables"`
	Text                string      `jsonapi:"attr,text"`
	StartDate           string      `jsonapi:"attr,startDate"`
	EndDate             string      `jsonapi:"attr,endDate"`
	SubjectUuid         string      `jsonapi:"attr,subject_uuid"`
	SubjectName         string      `jsonapi:"attr,subject_name"`
	TeacherUuid         string      `jsonapi:"attr,teacher_uuid"`
	TeacherName         string      `jsonapi:"attr,teacher_name"`
	RecurrenceRule      interface{} `jsonapi:"attr,recurrenceRule"`
	RecurrenceException interface{} `jsonapi:"attr,recurrenceException"`
	SessionsUuids       []string    `jsonapi:"attr,sessions_uuids"`
}

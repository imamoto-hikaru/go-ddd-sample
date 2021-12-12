package calendar

type CalendarRepository interface {
	Save(Name) (Calendar, error)
	Find(Id) (Calendar, error)
}

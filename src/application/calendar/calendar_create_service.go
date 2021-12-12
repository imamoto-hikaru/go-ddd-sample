package calendar

import "advent-calendar/domain/calendar"

type CalendarCreateService struct {
	calendarRepository calendar.CalendarRepository
}

func (service *CalendarCreateService) Create(name calendar.Name) (calendar.Calendar, error) {
	return service.calendarRepository.Save(name)
}

func NewCalendarCreateService(calendarRepository calendar.CalendarRepository) CalendarCreateService {
	return CalendarCreateService{calendarRepository: calendarRepository}
}

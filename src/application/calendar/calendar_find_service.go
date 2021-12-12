package calendar

import "advent-calendar/domain/calendar"

type CalendarFindService struct {
	calendarRepository calendar.CalendarRepository
}

func (service *CalendarFindService) Find(id calendar.Id) (calendar.Calendar, error) {
	return service.calendarRepository.Find(id)
}

func NewCalendarFindService(calendarRepository calendar.CalendarRepository) CalendarFindService {
	return CalendarFindService{calendarRepository: calendarRepository}
}

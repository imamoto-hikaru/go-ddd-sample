package entry

import (
	calendar_app "advent-calendar/application/calendar"
	"advent-calendar/domain/calendar"
	"advent-calendar/domain/entry"
)

type CalendarEntryListFindService struct {
	entryRepository     entry.EntryRepository
	calendarFindService calendar_app.CalendarFindService
}

func (service *CalendarEntryListFindService) Find(calendarId calendar.Id) (entry.CalendarEntryList, error) {
	calendar, err := service.calendarFindService.Find(calendarId)
	if err != nil {
		return entry.CalendarEntryList{}, err
	}

	return service.entryRepository.FindCalendarEntryList(calendar)
}

func NewCalendarEntryListFindService(
	entryRepository entry.EntryRepository,
	calendarFindService calendar_app.CalendarFindService) CalendarEntryListFindService {

	return CalendarEntryListFindService{
		entryRepository:     entryRepository,
		calendarFindService: calendarFindService,
	}
}

package entry

import (
	"advent-calendar/application/calendar"
	"advent-calendar/application/user"
	"advent-calendar/domain/entry"
	"errors"
)

type EntryCreateService struct {
	entryRepository     entry.EntryRepository
	userFindService     user.UserFindService
	calendarFindService calendar.CalendarFindService
}

func (service *EntryCreateService) Create(entrySheet entry.EntrySheet) error {
	if _, err := service.userFindService.Find(entrySheet.UserId); err != nil {
		return errors.New("user refer error:" + err.Error())
	}

	calendar, err := service.calendarFindService.Find(entrySheet.CalendarId)
	if err != nil {
		return errors.New("calendar refer error:" + err.Error())
	}

	calendarEntryList, err := service.entryRepository.FindCalendarEntryList(calendar)
	if err != nil {
		return errors.New("calendarEntryList refer error:" + err.Error())
	}

	if !calendarEntryList.CanAddEntryTo(entrySheet.EntriedDate) {
		return errors.New("already entried")
	}

	return service.entryRepository.Save(entrySheet)
}

func NewEntryCreateService(
	entryRepository entry.EntryRepository,
	userFindService user.UserFindService,
	calendarFindService calendar.CalendarFindService) EntryCreateService {
	return EntryCreateService{
		entryRepository:     entryRepository,
		userFindService:     userFindService,
		calendarFindService: calendarFindService,
	}
}

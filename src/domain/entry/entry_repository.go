package entry

import (
	"advent-calendar/domain/calendar"
	"advent-calendar/domain/user"
)

type EntryRepository interface {
	Save(EntrySheet) error
	FindCalendarEntryList(calendar.Calendar) (CalendarEntryList, error)
	FindUserEntryList(user.User) (UserEntryList, error)
}

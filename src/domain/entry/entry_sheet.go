package entry

import (
	"advent-calendar/domain/calendar"
	"advent-calendar/domain/user"
)

type EntrySheet struct {
	CalendarId  calendar.Id
	UserId      user.Id
	EntriedDate EntriedDate
}

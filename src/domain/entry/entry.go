package entry

import (
	"advent-calendar/domain/calendar"
	"advent-calendar/domain/user"
)

type Entry struct {
	Id          Id
	Calendar    calendar.Calendar
	User        user.User
	EntriedDate EntriedDate
}

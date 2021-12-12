package entry

import (
	"advent-calendar/domain/entry"
	"advent-calendar/infrastructure/calendar"
	"advent-calendar/infrastructure/user"
	"time"
)

type Entry struct {
	Id          uint
	CalendarId  uint
	UserId      uint
	EntriedDate uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Calendar    calendar.Calendar `gorm:"<-:false"`
	User        user.User         `gorm:"<-:false"`
}

func (m Entry) CreateEntity() entry.Entry {
	return entry.Entry{
		Id:          entry.Id(m.Id),
		Calendar:    m.Calendar.CreateEntity(),
		User:        m.User.CreateEntity(),
		EntriedDate: entry.EntriedDate(m.EntriedDate),
	}
}

package calendar

import (
	"advent-calendar/domain/calendar"
	"time"
)

type Calendar struct {
	Id           uint `gorm:"primaryKey"`
	CalendarName string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (m Calendar) CreateEntity() calendar.Calendar {
	return calendar.Calendar{Id: calendar.Id(m.Id), Name: calendar.Name(m.CalendarName)}
}

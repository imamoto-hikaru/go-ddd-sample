package calendar

import (
	"advent-calendar/domain/calendar"

	"gorm.io/gorm"
)

type CalendarRepository struct {
	db *gorm.DB
}

func (repository *CalendarRepository) Save(name calendar.Name) (calendar.Calendar, error) {
	m := Calendar{CalendarName: string(name)}
	result := repository.db.Create(&m)
	if err := result.Error; err != nil {
		return calendar.Calendar{}, err
	}

	return m.CreateEntity(), nil
}

func (repository *CalendarRepository) Find(id calendar.Id) (calendar.Calendar, error) {
	m := Calendar{}
	result := repository.db.First(&m, id)
	if err := result.Error; err != nil {
		return calendar.Calendar{}, err
	}

	return m.CreateEntity(), nil
}

func NewCalendarRepository(db *gorm.DB) *CalendarRepository {
	repository := new(CalendarRepository)
	repository.db = db
	return repository
}

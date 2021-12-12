package entry

import (
	"advent-calendar/domain/calendar"
	"advent-calendar/domain/entry"
	"advent-calendar/domain/user"

	"gorm.io/gorm"
)

type EntryRepository struct {
	db *gorm.DB
}

func (repository *EntryRepository) Save(entrySheet entry.EntrySheet) error {
	m := Entry{
		CalendarId:  uint(entrySheet.CalendarId),
		UserId:      uint(entrySheet.UserId),
		EntriedDate: uint(entrySheet.EntriedDate),
	}
	result := repository.db.Save(&m)
	return result.Error
}

func (repository *EntryRepository) FindCalendarEntryList(calendar calendar.Calendar) (entry.CalendarEntryList, error) {
	modelList := []Entry{}
	result := repository.db.Where(&Entry{CalendarId: uint(calendar.Id)}).Find(&modelList)
	if err := result.Error; err != nil {
		return entry.CalendarEntryList{}, err
	}

	entryList := []entry.Entry{}
	for _, e := range modelList {
		entryList = append(entryList, e.CreateEntity())
	}

	return entry.CalendarEntryList{
		Id:           calendar.Id,
		CalendarName: calendar.Name,
		EntryList:    entryList,
	}, nil
}

func (repository *EntryRepository) FindUserEntryList(user user.User) (entry.UserEntryList, error) {
	modelList := []Entry{}
	result := repository.db.Where(&Entry{UserId: uint(user.Id)}).Find(&modelList)
	if err := result.Error; err != nil {
		return entry.UserEntryList{}, err
	}

	entryList := []entry.Entry{}
	for _, e := range modelList {
		entryList = append(entryList, e.CreateEntity())
	}

	return entry.UserEntryList{
		Id:        user.Id,
		UserName:  user.Name,
		EntryList: entryList,
	}, nil
}

func NewEntryRepository(db *gorm.DB) *EntryRepository {
	repository := new(EntryRepository)
	repository.db = db
	return repository
}

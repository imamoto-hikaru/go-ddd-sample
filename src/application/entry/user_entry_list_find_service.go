package entry

import (
	user_app "advent-calendar/application/user"
	"advent-calendar/domain/entry"
	"advent-calendar/domain/user"
)

type UserEntryListFindService struct {
	entryRepository entry.EntryRepository
	userFindService user_app.UserFindService
}

func (service *UserEntryListFindService) Find(userId user.Id) (entry.UserEntryList, error) {
	user, err := service.userFindService.Find(userId)
	if err != nil {
		return entry.UserEntryList{}, err
	}

	return service.entryRepository.FindUserEntryList(user)
}

func NewUserEntryListFindService(
	entryRepository entry.EntryRepository,
	userFindService user_app.UserFindService) UserEntryListFindService {

	return UserEntryListFindService{
		entryRepository: entryRepository,
		userFindService: userFindService,
	}
}

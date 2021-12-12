package user

import "advent-calendar/domain/user"

type UserFindService struct {
	userRepository user.UserRepository
}

func (service *UserFindService) Find(id user.Id) (user.User, error) {
	return service.userRepository.Find(id)
}

func NewUserFindService(userRepository user.UserRepository) UserFindService {
	return UserFindService{userRepository: userRepository}
}

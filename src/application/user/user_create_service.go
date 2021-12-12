package user

import "advent-calendar/domain/user"

type UserCreateService struct {
	userRepository user.UserRepository
}

func (service *UserCreateService) Create(name user.Name) (user.User, error) {
	return service.userRepository.Save(name)
}

func NewUserCreateService(userRepository user.UserRepository) UserCreateService {
	return UserCreateService{userRepository: userRepository}
}

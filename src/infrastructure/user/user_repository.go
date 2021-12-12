package user

import (
	"advent-calendar/domain/user"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (repository *UserRepository) Save(name user.Name) (user.User, error) {
	m := User{UserName: string(name)}
	result := repository.db.Create(&m)
	if err := result.Error; err != nil {
		return user.User{}, err
	}
	return m.CreateEntity(), nil
}

func (repository *UserRepository) Find(id user.Id) (user.User, error) {
	m := User{}
	result := repository.db.First(&m, id)
	if err := result.Error; err != nil {
		return user.User{}, err
	}
	return m.CreateEntity(), nil
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	repository := new(UserRepository)
	repository.db = db
	return repository
}

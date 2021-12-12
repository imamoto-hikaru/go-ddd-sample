package user

import (
	"advent-calendar/domain/user"
	"time"
)

type User struct {
	Id        uint `gorm:"primaryKey"`
	UserName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m User) CreateEntity() user.User {
	return user.User{Id: user.Id(m.Id), Name: user.Name(m.UserName)}
}

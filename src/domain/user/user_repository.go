package user

type UserRepository interface {
	Save(Name) (User, error)
	Find(Id) (User, error)
}

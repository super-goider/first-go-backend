package users

type UserRepo interface {
	Add(u User) (User, error)
	GetByID(id int) (User, bool, error)
	GetByLogin(login string) (User, bool, error)
	All() ([]User, error)
}

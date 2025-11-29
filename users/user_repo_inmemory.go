package users

type InMemoryUserRepo struct {
	users []User
}

func NewInMemory() *InMemoryUserRepo {
	return &InMemoryUserRepo{users: make([]User, 0)}
}

func (mem *InMemoryUserRepo) Add(u User) (User, error) {
	u.ID = len(mem.users) + 1
	mem.users = append(mem.users, u)
	return u, nil
}

func (mem *InMemoryUserRepo) GetByID(id int) (User, bool, error) {
	for _, u := range mem.users {
		if u.ID == id {
			return u, true, nil
		}
	}
	return User{}, false, nil
}
func (mem *InMemoryUserRepo) GetByLogin(login string) (User, bool, error) {
	for _, u := range mem.users {
		if u.Login == login {
			return u, true, nil
		}
	}
	return User{}, false, nil
}

func (mem *InMemoryUserRepo) All() ([]User, error) {
	usersCopy := make([]User, len(mem.users))
	copy(usersCopy, mem.users)
	return usersCopy, nil
}

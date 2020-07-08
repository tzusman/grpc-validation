package util

type User struct {
	Name        string
	Age         int32
	Email       string
	PhoneNumber string
	Attributes  []string
}

type UserStore struct {
	store map[string]User
}

func NewUserStore() UserStore {
	return UserStore{
		store: make(map[string]User),
	}
}

func (s UserStore) AddUser(id string, user User) {
	s.store[id] = user
}

func (s UserStore) RetrieveUser(id string) User {
	return s.store[id]
}

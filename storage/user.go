package storage

import (
	"errors"
	"github.com/rthornton128/test/object"
)

// UserStorage a storage interface for Users
type UserStorage interface {
	Get(u *object.User) error
	GetAll() []object.User
}

type userStorage struct {}

// NewUserStorage returns a database storage for Users
func NewUserStorage(db string) UserStorage {
	return &userStorage{}
}

func (us *userStorage) Get(u *object.User) error {
	switch u.ID {
	case 1:
		u.First = "Admin"
		u.Last = "User"
	case 2:
		u.First = "Bob"
		u.Last = "Cajun"
	case 3:
		u.First = "Frank"
		u.Last = "Furter"
	default:
		return errors.New("no user found")
	}
	return nil
}

func (us *userStorage) GetAll() []object.User {
	var users []object.User
	users = append(users, object.User{ID: 1, First: "Admin", Last: "User"})
	users = append(users, object.User{ID: 2, First: "Bob", Last: "Cajun"})
	users = append(users, object.User{ID: 3, First: "Frank", Last: "Furter"})
	return users
}

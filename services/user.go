package services

import (
	"encoding/json"
	"errors"
	"github.com/rthornton128/test/object"
	"github.com/rthornton128/test/storage"
)

// UserService interface
type UserService interface {
	Get(ID int64) (string, error)
	GetAll() (string, error)
}

// User object
type userService struct {
	Storage storage.UserStorage 
}

// NewUserService creates a new user object
func NewUserService(us storage.UserStorage) UserService {
	return &userService{us}
}

// Get User JSON string
func (s *userService) Get(ID int64) (string, error) {
	u := object.User{ID: ID}
	err := s.Storage.Get(&u)
	if err != nil {
		return "", errors.New("userService.Get: " + err.Error())
	}
	jsonString, err := json.Marshal(&u)
	if err != nil {
		return "", errors.New("Marshalling user JSON string: " + err.Error())
	}
	return string(jsonString), nil
}

func (s *userService) GetAll() (string, error) {
	users := s.Storage.GetAll()
	jsonString, err := json.Marshal(&users)
	if err != nil {
		return "", errors.New("Marshalling user JSON string: " + err.Error())
	}
	return string(jsonString), nil

}
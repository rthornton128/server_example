package services_test

import (
	"errors"
	"testing"

	"github.com/rthornton128/test/object"
	"github.com/rthornton128/test/services"
)

var userData = []object.User{
	{ID: 1, First: "User", Last: "Name"},
}

type mockUserStorage struct{}

func NewUserStorage() *mockUserStorage {
	return &mockUserStorage{}
}

func (m *mockUserStorage) Get(u *object.User) error {
	if u == nil {
		return errors.New("nil user object")
	}
	if u.ID != 1 {
		return errors.New("unknown ID")
	}
	u.ID = userData[0].ID
	u.First = userData[0].First
	u.Last = userData[0].Last
	return nil
}

func (m *mockUserStorage) GetAll() []object.User {
	return userData
}

func TestUserServiceGet(t *testing.T) {
	storage := NewUserStorage()
	userService := services.NewUserService(storage)
	str, err := userService.Get(1)
	t.Log(str)
	if err != nil {
		t.Fatal("service error: ", err)
	}

	str, err = userService.Get(2)
	t.Log(str)
	if err == nil {
		t.Fatal("expected call to err to be nil, got: ", err)
	}
}

func TestUserServiceGetAll(t *testing.T) {
	storage := NewUserStorage()
	userService := services.NewUserService(storage)
	_, err := userService.GetAll()
	if err != nil {
		t.Fatal("expected err to be nil, got: ", err)
	}
}

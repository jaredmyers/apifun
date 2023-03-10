package services

import (
	"log"

	"github.com/jaredmyers/apifun/go_api/models"
	"github.com/jaredmyers/apifun/go_api/storage"
)

type UserService struct {
	store storage.UserServiceStorer
}

func NewUserService(store storage.UserServiceStorer) UserServicer {
	return &UserService{
		store: store,
	}
}

func (uc *UserService) CreateUser(*models.User) error {
	return nil
}
func (uc *UserService) GetUser(int) (*models.User, error) {

	return nil, nil
}
func (uc *UserService) UpdateUser(*models.User) error {
	return nil
}
func (db *UserService) DeleteUser(*string) error {
	return nil
}
func (uc *UserService) GetUsers() ([]*models.User, error) {
	users, err := uc.store.GetUsers()
	if err != nil {
		return nil, err
	}
	log.Println("from UserService GetUsers")
	log.Println(users)
	return users, nil
}

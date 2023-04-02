package services

import (
	"context"

	"github.com/jaredmyers/apifun/go_api/models"
	"github.com/jaredmyers/apifun/go_api/storage"
)

type UserService struct {
	store storage.UserServiceStorer
	cache storage.UserServiceCacher
}

func NewUserService(store storage.UserServiceStorer, cache storage.UserServiceCacher) UserServicer {
	return &UserService{
		store: store,
		cache: cache,
	}
}

func (uc *UserService) CreateUser(*models.User) error {
	return nil
}
func (uc *UserService) GetUser(userId int) (*models.User, error) {

	// cache check
	user, err := uc.cache.GetUser(context.Background(), userId)
	if err == nil {
		return user, nil
	}

	// if cache miss, go to database
	user, err = uc.store.GetUser(userId)
	if err != nil {
		return nil, err
	}

	// store value in cache
	if err := uc.cache.SetUser(context.Background(), user); err != nil {
		return nil, err
	}

	return user, nil
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
	return users, nil
}

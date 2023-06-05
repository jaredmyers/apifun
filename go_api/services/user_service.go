package services

import (
	"context"
	"log"

	m "github.com/jaredmyers/apifun/go_api/models"
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

func (uc *UserService) RegisterUser(req *m.RegisterUserRequest) error {
	log.Println("RegisterUser from userservice running...")
	log.Println(req)

	err := uc.store.RegisterUser(req)
	if err != nil {
		return err
	}

	return nil
}
func (uc *UserService) GetUser(userId int) (*m.User, error) {

	log.Println("GetUser from UserService running...")

	// cache check if a cache has been created

	if uc.cache != nil {
		user, err := uc.cache.GetUser(context.Background(), userId)
		if err == nil {
			return user, nil
		}
	}

	// if cache miss, go to database
	user, err := uc.store.GetUser(userId)
	if err != nil {
		return nil, err
	}

	// set value in cache if cache has been created
	if uc.cache != nil {
		if err := uc.cache.SetUser(context.Background(), user); err != nil {
			return nil, err
		}
	}

	return user, nil
}
func (uc *UserService) UpdateUser(*m.User) error {
	return nil
}
func (db *UserService) DeleteUser(*string) error {
	return nil
}
func (uc *UserService) GetUsers() ([]*m.User, error) {
	users, err := uc.store.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

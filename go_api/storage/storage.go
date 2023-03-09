package storage

import "github.com/jaredmyers/apifun/go_api/models"

type UserServicer interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
	GetUsers() error
}

type FoodServicer interface {
}

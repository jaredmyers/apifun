package services

import "github.com/jaredmyers/apifun/go_api/models"

type UserServicer interface {
	CreateUser(*models.User) error
	GetUser(int) (*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
	GetUsers() ([]*models.User, error)
}

type FoodServicer interface {
}

package storage

import "github.com/jaredmyers/apifun/go_api/models"

type UserServiceStorer interface {
	CreateUser(*models.User) error
	GetUser(int) (*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(int) error
	GetUsers() ([]*models.User, error)
}

package services

import (
	m "github.com/jaredmyers/apifun/go_api/models"
)

type UserServicer interface {
	RegisterUser(*m.RegisterUserRequest) error
	GetUser(int) (*m.User, error)
	UpdateUser(*m.User) error
	DeleteUser(*string) error
	GetUsers() ([]*m.User, error)
}

type FoodServicer interface {
}

package storage

import (
	"context"

	m "github.com/jaredmyers/apifun/go_api/models"
)

type UserServiceStorer interface {
	CreateUser(*m.User) error
	GetUser(int) (*m.User, error)
	UpdateUser(*m.User) error
	DeleteUser(int) error
	GetUsers() ([]*m.User, error)
}

type UserServiceCacher interface {
	GetUser(context.Context, int) (*m.User, error)
	SetUser(context.Context, *m.User) error
}

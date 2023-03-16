package storage

import (
	"github.com/jaredmyers/apifun/go_api/models"
)

type MockStore struct{}

func NewMockStore() (UserServiceStorer, error) {
	return &MockStore{}, nil
}

func (m *MockStore) CreateUser(*models.User) error {
	return nil
}
func (m *MockStore) GetUser(*string) (*models.User, error) {
	return nil, nil
}
func (m *MockStore) UpdateUser(*models.User) error {
	return nil
}
func (m *MockStore) DeleteUser(*string) error {
	return nil
}
func (m *MockStore) GetUsers() ([]*models.User, error) {
	// select * from users;

	users := []*models.User{
		{Id: 0, Username: "FoxTrot", Pw: "123", CreatedOn: "Date", DeletedOn: nil},
		{Id: 1, Username: "Tango", Pw: "123", CreatedOn: "Date", DeletedOn: nil},
		{Id: 2, Username: "Charlie", Pw: "123", CreatedOn: "Date", DeletedOn: nil},
		{Id: 3, Username: "Aplha", Pw: "123", CreatedOn: "Date", DeletedOn: nil},
		{Id: 4, Username: "Midnight", Pw: "123", CreatedOn: "Date", DeletedOn: nil},
		{Id: 5, Username: "Bravo", Pw: "123", CreatedOn: "Date", DeletedOn: nil},
	}

	return users, nil
}

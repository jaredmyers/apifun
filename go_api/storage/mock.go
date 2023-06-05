package storage

import (
	m "github.com/jaredmyers/apifun/go_api/models"
)

type MockStore struct{}

func NewMockStore() (UserServiceStorer, error) {
	return &MockStore{}, nil
}

func (mock *MockStore) RegisterUser(*m.RegisterUserRequest) error {
	return nil
}
func (mock *MockStore) GetUser(id int) (*m.User, error) {
	return nil, nil
}
func (mock *MockStore) UpdateUser(*m.User) error {
	return nil
}
func (mock *MockStore) DeleteUser(id int) error {
	return nil
}
func (mock *MockStore) GetUsers() ([]*m.User, error) {
	// select * from users;

	users := []*m.User{
		{Id: 0, Username: "FoxTrot", Pw: "123", CreatedOn: "Date", DeletedOn: nil},
		{Id: 1, Username: "Tango", Pw: "123", CreatedOn: "Date", DeletedOn: nil},
		{Id: 2, Username: "Charlie", Pw: "123", CreatedOn: "Date", DeletedOn: nil},
		{Id: 3, Username: "Aplha", Pw: "123", CreatedOn: "Date", DeletedOn: nil},
		{Id: 4, Username: "Midnight", Pw: "123", CreatedOn: "Date", DeletedOn: nil},
		{Id: 5, Username: "Bravo", Pw: "123", CreatedOn: "Date", DeletedOn: nil},
	}

	return users, nil
}

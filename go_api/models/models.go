package models

// User is a user entity of the system
type User struct {
	Id        int     `json:"id"`
	Username  string  `json:"username"`
	Pw        string  `json:"pw"`
	CreatedOn string  `json:"createdon"`
	DeletedOn *string `json:"deletedon,omitempty"`
}

// GetUsersResponse defines response returned back after searching/getting users
type GetUsersResponse struct {
	Users []*User `json:"users"`
}

type GetUserRequest int
type LoginRequest struct{}
type CreateUserRequest struct{}

type ClosestBusResponse struct {
	Id int
}

package models

// User is a user entity of the system
type User struct {
	Id        int     `json:"id"`
	Username  string  `json:"username"`
	Pw        string  `json:"pw"`
	CreatedOn string  `json:"createdon"`
	DeletedOn *string `json:"deletedon,omitempty"`
}

// requests
type RegisterUserRequest struct {
	Username string `json:"username"`
	Pw       string `json:"pw"`
}

type GetUserRequest int
type LoginRequest struct{}

// responses
type GetUsersResponse struct {
	Users []*User `json:"users"`
}
type ClosestBusResponse struct {
	Id int
}

type RegisterUserResponse struct {
	Status string `json:"status"`
}

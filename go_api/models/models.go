package models

type User struct {
	Id        int     `json:"id"`
	Username  string  `json:"username"`
	Pw        string  `json:"pw"`
	CreatedOn string  `json:"createdon"`
	DeletedOn *string `json:"deletedon,omitempty"`
}

type GetUsersResponse struct {
	Users []*User `json:"users"`
}

type getUserRequest int

type LoginRequest struct {
}
type CreateUserRequest struct {
}

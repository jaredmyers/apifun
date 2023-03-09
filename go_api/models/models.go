package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Pw       string `json:"pw"`
}

type LoginRequest struct {
}
type CreateUserRequest struct {
}

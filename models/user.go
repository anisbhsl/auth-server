package models

type User struct {
	ID           string `json:"id"`
	Email        string `json:"email" validate:"nonzero"`
	Name         string `json:"name"`
	Location     string `json:"location"`
	About        string `json:"about"`
	PasswordHash string `json:"password_hash"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"nonzero"`
	Password string `json:"password" validate:"nonzero"`
}

type RegisterUserRequest struct {
	User
	Password string `json:"password" validate:"nonzero"`
}

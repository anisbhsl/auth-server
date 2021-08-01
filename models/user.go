package models

type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Location     string `json:"location"`
	About        string `json:"about"`
	PasswordHash string `json:"password_hash"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserRequest struct {
	User
	Password string `json:"password"`
}

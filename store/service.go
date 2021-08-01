package store

import (
	"github.com/anisbhsl/auth-server/models"
)

type Service interface {
	GetUserDetail(id string) (models.User,error)
	GetUserByEmail(email string)(models.User,error)
	AddUser(user models.User) (string,error)
}


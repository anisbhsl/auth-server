package mock

import (
	"github.com/anisbhsl/auth-server/models"
)

type Store struct{
	User models.User
}

func (s Store) GetUserDetail(id string) (models.User,error){
	return s.User,nil
}

func (s Store) GetUserByEmail(email string)(models.User,error){
	return s.User,nil
}

func (s Store) AddUser(user models.User) (string,error){
	s.User=user
	return user.ID,nil
}
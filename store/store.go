package store

import "github.com/anisbhsl/auth-server/models"

type Service interface {
	GetUserDetail(id string) (models.User,error)
	GetUserByEmail(email string)(models.User,error)
	AddUser(user models.User) (string,error)
}

type store struct{

}

func New() Service{
	return &store{}
}

func (s store) GetUserDetail(id string) (models.User,error){
	return models.User{},nil
}

func (s store) 	GetUserByEmail(email string)(models.User,error){
	return models.User{},nil
}

func (s store) AddUser(user models.User) (string,error){
	return user.ID,nil
}
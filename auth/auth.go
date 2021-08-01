package auth

import (
	"fmt"
	"github.com/anisbhsl/auth-server/models"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GenerateToken() (models.Token, error)
	VerifyToken()
	TokenRefresh(token string) (models.Token, error)
	EncryptPassword(string) string
	ValidatePasswordHash(raw string, hash string) bool
}

type service struct {
	Secret string
}

func New(secretKey string) Service {
	return &service{Secret: secretKey}
}

func (s service) VerifyToken() {

}

func (s service) TokenRefresh(token string) (models.Token, error) {
	return models.Token{}, nil
}

func (s service) GenerateToken() (models.Token, error) {
	return models.Token{}, nil
}

//EncryptPassword encrypts given pass with secret as salt
func (s service) EncryptPassword(src string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(s.getRawPasswordWithSalt(src)), bcrypt.MinCost)
	return string(bytes)
}


func (s service) ValidatePasswordHash(raw string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(s.getRawPasswordWithSalt(raw)))
	return err==nil
}

func (s service) getRawPasswordWithSalt(src string)string{
	return fmt.Sprintf("%s:%s",s.Secret,src)
}

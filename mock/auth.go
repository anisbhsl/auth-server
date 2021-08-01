package mock

import (
	"github.com/anisbhsl/auth-server/models"
	"github.com/golang-jwt/jwt"
)

type AuthService struct{
	AccessToken string
	RefreshToken string
}

func (a AuthService) GenerateToken(email string) (models.Token, error){
	return models.Token{
		AccessToken:  a.AccessToken,
		RefreshToken: a.RefreshToken,
	},nil
}

func (a AuthService) VerifyToken(signedJWT string) (*models.UserClaims,error){
	return &models.UserClaims{
		StandardClaims: jwt.StandardClaims{},
		Email:  "test@test.com",
	},nil
}

func (a AuthService) TokenRefresh(token string) (models.Token, error){
	return models.Token{
		AccessToken:  a.AccessToken,
		RefreshToken: a.RefreshToken,
	},nil
}

func (a AuthService) EncryptPassword(src string) string{
	return src
}

func (a AuthService) ValidatePasswordHash(raw string, hash string) bool{
	return raw==hash
}

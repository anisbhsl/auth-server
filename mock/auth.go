package mock

import (
	"github.com/anisbhsl/auth-server/auth"
	"github.com/anisbhsl/auth-server/models"
	"github.com/square/go-jose/v3/jwt"
)

type AuthService struct{

}

func (a AuthService) GenerateToken(email string) (models.Token, error){
	return models.Token{
		AccessToken:  "213",
		RefreshToken: "123",
	},nil
}

func (a AuthService) VerifyToken(signedJWT string) (*auth.UserClaims,error){
	return &auth.UserClaims{
		Claims: jwt.Claims{},
		Email:  "test@test.com",
	},nil
}

func (a AuthService) TokenRefresh(token string) (models.Token, error){
	return models.Token{
		AccessToken:  "213",
		RefreshToken: "123",
	},nil
}

func (a AuthService) EncryptPassword(src string) string{
	return src
}

func (a AuthService) ValidatePasswordHash(raw string, hash string) bool{
	return true
}

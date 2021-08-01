package auth

import (
	"github.com/anisbhsl/auth-server/models"
)

type Service interface {
	GenerateToken(email string) (models.Token, error)
	VerifyToken(signedJWT string) (*models.UserClaims,error)
	TokenRefresh(token string) (models.Token, error)
	EncryptPassword(string) string
	ValidatePasswordHash(raw string, hash string) bool
}

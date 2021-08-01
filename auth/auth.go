package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"github.com/anisbhsl/auth-server/logger"
	"github.com/anisbhsl/auth-server/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/square/go-jose/v3"
	"github.com/square/go-jose/v3/jwt"
	"time"
)

const(
	Issuer="XYZ Corp"
	AccessTokenExpiryDuration=1*time.Hour
	RefreshTokenExpiryDuration=3*time.Hour
)

type UserClaims struct{
	jwt.Claims
	Email string `json:"email"`
}

type service struct {
	Secret string
	PrivateKey *rsa.PrivateKey
	PublicKey *rsa.PublicKey
}

func New(secretKey string) Service {
	return &service{Secret: secretKey}
}

func (s service) VerifyToken(signedJWT string) (*UserClaims,error) {
	token,err:=jwt.ParseSigned(signedJWT)
	if err!=nil{
		return nil,err
	}

	claims:=&UserClaims{}
	if err:=token.Claims(s.PublicKey,claims);err!=nil{
		return nil,err
	}

	err=claims.Validate(jwt.Expected{
		Issuer:   Issuer,
		Time:     time.Now(),
	})
	if err!=nil{
		if err==jwt.ErrExpired{
			return nil,fmt.Errorf("invalid token")
		}
		return nil, fmt.Errorf("invalid token")
	}
	return claims,nil
}

func (s service) TokenRefresh(token string) (models.Token, error) {
	claims,err:=s.VerifyToken(token)
	if err!=nil{
		logger.Error(fmt.Sprintf("%v",err))
		return models.Token{},fmt.Errorf("Invalid Token")
	}
	signedTokens,err:=s.GenerateToken(claims.Email)
	if err!=nil{
		return models.Token{},nil
	}
	return signedTokens, nil
}

func NewRSAKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}

func (s service) GenerateToken(email string) (models.Token, error) {
	key,_:=NewRSAKey()
	//key,_:=rsa.GenerateKey()
	accessTokenClaims:=UserClaims{
		Claims: jwt.Claims{
			Issuer:Issuer,
			Expiry: jwt.NewNumericDate(time.Now().Add(AccessTokenExpiryDuration)),
		},
		Email:  email,
	}

	refreshTokenClaims:=UserClaims{
		Claims: jwt.Claims{
			Issuer:Issuer,
			Expiry: jwt.NewNumericDate(time.Now().Add(RefreshTokenExpiryDuration)),
		},
		Email:  email,
	}

	accessToken,err:=s.getToken(accessTokenClaims,key)
	if err!=nil{
		logger.Error(fmt.Sprintf("%v",err))
		return models.Token{},err
	}
	refreshToken,err:=s.getToken(refreshTokenClaims,key)
	if err!=nil{
		logger.Error(fmt.Sprintf("%v",err))
		return models.Token{},err
	}

	return models.Token{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}, nil
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

func (s service) getToken(claims UserClaims, key interface{}) (string,error){
	opts:=jose.SignerOptions{}
	opts.WithType("JWT")

	signer,err:=jose.NewSigner(jose.SigningKey{
		Algorithm: jose.RS256,
		Key:       key,
	},&opts)
	if err!=nil{
		return "",err
	}
	return jwt.Signed(signer).Claims(claims).CompactSerialize()
}

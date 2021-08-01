package auth

import (
	"crypto/rsa"
	"fmt"
	"github.com/anisbhsl/auth-server/logger"
	"github.com/anisbhsl/auth-server/models"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"time"
)

const(
	Issuer="XYZ Corp"
	AccessTokenExpiryDuration=1*time.Hour
	RefreshTokenExpiryDuration=3*time.Hour
)

var tokenType=struct{
	Access string
	Refresh string
}{
	"access",
	"refresh",
}

type service struct {
	Secret string
	PrivateKey *rsa.PrivateKey
	PublicKey *rsa.PublicKey
}

func New(secretKey, privateKeyPath, publicKeyPath string) Service {
	pvtBytes,err:=ioutil.ReadFile(privateKeyPath)
	if err!=nil{
		panic(err)
	}

	privateKey,err:=jwt.ParseRSAPrivateKeyFromPEM(pvtBytes)
	if err!=nil{
		panic(err)
	}

	publicBytes,err:=ioutil.ReadFile(publicKeyPath)
	if err!=nil{
		panic(err)
	}

	publicKey,err:=jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err!=nil{
		panic(err)
	}

	return &service{Secret: secretKey,PrivateKey: privateKey,PublicKey: publicKey}
}

func (s service) GenerateToken(email string) (models.Token, error) {
	accessTokenClaims:=models.UserClaims{
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccessTokenExpiryDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    Issuer,
		},
		Email:  email,
		TokenType: tokenType.Access,
	}

	refreshTokenClaims:=models.UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(RefreshTokenExpiryDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    Issuer,
		},
		Email:  email,
		TokenType: tokenType.Refresh,
	}

	accessToken,err:=s.getToken(accessTokenClaims)
	if err!=nil{
		logger.Error(fmt.Sprintf("%v",err))
		return models.Token{},err
	}
	refreshToken,err:=s.getToken(refreshTokenClaims)
	if err!=nil{
		logger.Error(fmt.Sprintf("%v",err))
		return models.Token{},err
	}

	return models.Token{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}, nil
}


func (s service) VerifyToken(signedJWT string) (*models.UserClaims,error) {
	claims:=&models.UserClaims{}
	token,err:=jwt.ParseWithClaims(signedJWT,claims,func(j *jwt.Token)(interface{},error){
		return s.PublicKey,nil
	})
	if err!=nil{
		return nil, err
	}

	if claims,ok:=token.Claims.(*models.UserClaims);ok && token.Valid{
		return claims,nil
	}

	return nil,fmt.Errorf("Invalid Token")
}

func (s service) TokenRefresh(token string) (models.Token, error) {
	claims,err:=s.VerifyToken(token)
	if err!=nil || claims.TokenType==tokenType.Access{
		logger.Error(fmt.Sprintf("%v",err))
		return models.Token{},fmt.Errorf("Invalid Token")
	}

	signedTokens,err:=s.GenerateToken(claims.Email)
	if err!=nil{
		return models.Token{},nil
	}
	return signedTokens, nil
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

func (s service) getToken(claims models.UserClaims) (string,error){
	signer:=jwt.New(jwt.GetSigningMethod("RS256"))
	signer.Claims=claims
	return signer.SignedString(s.PrivateKey)
}

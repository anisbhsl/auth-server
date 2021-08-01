package models

import "github.com/golang-jwt/jwt"

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserClaims struct{
	jwt.StandardClaims
	TokenType string `json:"type"`
	Email string `json:"email"`
}
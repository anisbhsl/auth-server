package api

import (
	"encoding/json"
	"fmt"
	"github.com/anisbhsl/auth-server/logger"
	"github.com/anisbhsl/auth-server/models"
	"net/http"
	"gopkg.in/validator.v2"
)

func (s service) Login() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var loginRequest models.LoginRequest
		err:=json.NewDecoder(r.Body).Decode(&loginRequest)
		if err!=nil{
			logger.Error(fmt.Sprintf("%v",err),logger.TraceRequestWithContext(r.Context()))
			SendErrorResponse(w,CodeInvalidLoginCredentials)
			return
		}

		if err:=validator.Validate(loginRequest);err!=nil{
			SendErrorResponse(w,fmt.Sprintf("%v",err))
			return
		}

		user,err:=s.Store.GetUserByEmail(loginRequest.Email)
		if err!=nil{
			errMsg:=fmt.Sprintf("%v",err)
			logger.Error(errMsg,logger.TraceRequestWithContext(r.Context()))
			SendErrorResponse(w, errMsg)
			return
		}

		if !s.AuthService.ValidatePasswordHash(loginRequest.Password,user.PasswordHash){
			SendErrorResponse(w,CodeInvalidLoginCredentials)
			return
		}

		token,err:=s.AuthService.GenerateToken(user.Email)
		if err!=nil{
			logger.Error(fmt.Sprintf("%v",err),logger.TraceRequestWithContext(r.Context()))
			SendErrorResponse(w,CodeErrorTokenGeneration)
			return
		}

		SendSuccessResponse(w,models.Token{
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
		})
	}
}

func (s service) RefreshToken() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var refreshToken models.Token
		err:=json.NewDecoder(r.Body).Decode(&refreshToken)
		if err!=nil{
			logger.Error(fmt.Sprintf("%v",logger.TraceRequestWithContext(r.Context())))
			SendErrorResponse(w,CodeInvalidRegistrationData)
			return
		}

		if err:=validator.Validate(refreshToken);err!=nil{
			SendErrorResponse(w,fmt.Sprintf("%v",err))
			return
		}

		token,err:=s.AuthService.TokenRefresh(refreshToken.RefreshToken)
		if err!=nil{
			SendUnauthorizedResponse(w)
			return
		}

		SendSuccessResponse(w,models.Token{
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
		})

	}
}

package api

import (
	"encoding/json"
	"fmt"
	"github.com/anisbhsl/auth-server/logger"
	"github.com/anisbhsl/auth-server/models"
	"net/http"
)

func (s service) GetMeUser() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		userID,ok:=r.Context().Value("identity").(string)
		if !ok{
			SendErrorResponse(w,CodeInvalidToken)
			return
		}

		user,err:=s.Store.GetUserDetail(userID)
		if err!=nil{
			errMsg:=fmt.Sprintf("%v",err)
			logger.Error(errMsg,logger.TraceRequestWithContext(r.Context()))
			SendErrorResponse(w,errMsg)
			return
		}

		SendSuccessResponse(w,map[string]interface{}{
			"id":user.ID,
			"email":user.Email,
			"location":user.Location,
			"about":user.About,
		})
	}
}


func (s service) RegisterUser() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var registerUserRequest models.RegisterUserRequest
		err:=json.NewDecoder(r.Body).Decode(&registerUserRequest)
		if err!=nil{
			logger.Error(fmt.Sprintf("%v",err),logger.TraceRequestWithContext(r.Context()))
			SendErrorResponse(w, CodeInvalidRegistrationData)
			return
		}

		user:=models.User{
			ID:       s.IDGenerator.UserId(),
			Email:    registerUserRequest.Email,
			Name:     registerUserRequest.Name,
			Location: registerUserRequest.Location,
			About:    registerUserRequest.About,
			PasswordHash: s.AuthService.EncryptPassword(registerUserRequest.Password),
		}

		id,err:=s.Store.AddUser(user)
		if err!=nil{
			errMsg:=fmt.Sprintf("%v",err)
			logger.Error(errMsg,logger.TraceRequestWithContext(r.Context()))
			SendErrorResponse(w,errMsg)
		}
		SendSuccessResponse(w,map[string]interface{}{
			"id":id,
		})
	}
}
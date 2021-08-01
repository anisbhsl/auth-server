package api

import (
	"context"
	"github.com/anisbhsl/auth-server/mock"
	"github.com/anisbhsl/auth-server/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetMeUser(t *testing.T) {
	store:=mock.Store{User: models.User{
		ID:           "user_3lKRPAqj-AHFgjvkH3L_4",
		Email:        "bhusal.anish12@gmail.com",
		Name:         "Anish Bhusal",
		Location:     "Kathmandu",
		About:        "I am software dev!",
	}}
	service:=New(mock.AuthService{},store)
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/me",service.GetMeUser()).Methods("GET")
	out := httptest.NewRecorder()
	in:=httptest.NewRequest("GET","/api/v1/me",nil)
	in.Header.Add("Authorization","Bearer access_123")
	r.ServeHTTP(out,in.WithContext(context.WithValue(in.Context(),"email",store.User.Email)))
	assert.Equal(t,200,out.Code)
	assert.JSONEq(t,`{"data":{"about":"I am software dev!","email":"bhusal.anish12@gmail.com","id":"user_3lKRPAqj-AHFgjvkH3L_4","location":"Kathmandu"},"success":true}`,out.Body.String())
}

func TestRegisterUser(t *testing.T) {
	service:=New(mock.AuthService{},mock.Store{})
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/register-user",service.RegisterUser()).Methods("POST")
	out := httptest.NewRecorder()
	body := `{
	"name":"Anish Bhusal",
	"email":"bhusal.anish@gmail.com",
	"location":"Kathmandu",
	"about":"I am a software dev!!",
	"password":"letmeinplease"
	}`
	in := httptest.NewRequest("POST", "/api/v1/register-user", strings.NewReader(body))
	r.ServeHTTP(out, in)
	assert.Equal(t, 200, out.Code)
}

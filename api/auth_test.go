package api

import (
	"encoding/json"
	"github.com/anisbhsl/auth-server/mock"
	"github.com/anisbhsl/auth-server/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	store:=mock.Store{User: models.User{
		ID:           "user_3lKRPAqj-AHFgjvkH3L_4",
		Email:        "bhusal.anish12@gmail.com",
		Name:         "Anish Bhusal",
		Location:     "Kathmandu",
		About:        "I am software dev!",
		PasswordHash: "letmeinplease",
	}}
	service := New(mock.AuthService{AccessToken: "213",RefreshToken: "123"},store)
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/auth/login", service.Login()).Methods("POST")
	out := httptest.NewRecorder()
	bodyInBytes,_:=json.Marshal(models.LoginRequest{
		Email:    "bhusal.anish12@gmail.com",
		Password: "letmeinplease",
	})

	in := httptest.NewRequest("POST", "/api/v1/auth/login", strings.NewReader(string(bodyInBytes)))
	r.ServeHTTP(out, in)
	assert.Equal(t,200,out.Code)
	assert.JSONEq(t,`{"data":{"access_token":"213","refresh_token":"123"},"success":true}`,out.Body.String())
}

func TestRefreshToken(t *testing.T){
	service := New(mock.AuthService{AccessToken: "213",RefreshToken: "123"},mock.Store{})
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/auth/refresh-token", service.RefreshToken()).Methods("POST")
	out := httptest.NewRecorder()
	bodyInBytes,_ :=json.Marshal(models.Token{
		RefreshToken: "123",
	})
	in := httptest.NewRequest("POST", "/api/v1/auth/refresh-token", strings.NewReader(string(bodyInBytes)))
	r.ServeHTTP(out, in)
	assert.Equal(t,200,out.Code)
	assert.JSONEq(t,`{"data":{"access_token":"213","refresh_token":"123"},"success":true}`,out.Body.String())
}

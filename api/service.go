package api

import (
	"github.com/anisbhsl/auth-server/auth"
	"github.com/anisbhsl/auth-server/provider"
	"github.com/anisbhsl/auth-server/store"
	"net/http"
)

type Service interface {
	Index() http.HandlerFunc
	Login() http.HandlerFunc
	RefreshToken() http.HandlerFunc
	GetMeUser() http.HandlerFunc
	RegisterUser() http.HandlerFunc
	LoggingMiddleware(next http.Handler) http.Handler
	AuthMiddleware(next http.Handler) http.Handler
}

type service struct {
	AuthService auth.Service
	Store       store.Service
	IDGenerator provider.IDGenerator
}

func New(auth auth.Service, store store.Service) Service {
	return &service{
		AuthService: auth,
		Store:       store,
		IDGenerator: provider.NanoIDGenerator{},
	}
}

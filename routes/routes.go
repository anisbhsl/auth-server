package routes

import (
	"github.com/anisbhsl/auth-server/api"
	"github.com/anisbhsl/auth-server/middlewares"
	"github.com/anisbhsl/auth-server/utils"
	"github.com/gorilla/mux"
	chain "github.com/justinas/alice"
	"net/http"
)

var(
	loggingMiddleware=middlewares.LoggingMiddleware
	authMiddleware=middlewares.AuthMiddleware
)

var httpMethods = struct {
	GET    string
	POST   string
	DELETE string
	PUT    string
	PATCH  string
}{
	"GET",
	"POST",
	"DELETE",
	"PUT",
	"PATCH",
}

type routeConfig struct {
	Handler     http.Handler
	Methods     []string
	Middlewares []chain.Constructor
}

var routes = func(api api.Service) map[string]routeConfig {
	GENERAL := map[string]routeConfig{
		"/": {
			Handler:     api.Index(),
			Methods:     []string{httpMethods.GET},
			Middlewares: nil,
		},
	}

	AUTH := map[string]routeConfig{
		"/auth/login": {
			Handler: api.Login(),
			Methods: []string{httpMethods.POST},
		},
		"/auth/refresh-token": {
			Handler: api.RefreshToken(),
			Methods: []string{httpMethods.POST},
		},
	}

	USER:=map[string]routeConfig{
		"/me":{
			Handler:     api.GetMeUser(),
			Methods:     []string{httpMethods.GET},
			Middlewares: []chain.Constructor{authMiddleware},
		},
		"/register-user":{
			Handler: api.RegisterUser(),
			Methods: []string{httpMethods.POST},
		},
	}

	return func(routeMaps ...map[string]routeConfig) map[string]routeConfig {
		routeDefs := routeMaps[0]
		for i := 1; i < len(routeMaps); i++ {
			for path, config := range routeMaps[i] {
				routeDefs[path] = config
			}
		}
		return routeDefs
	}(GENERAL, AUTH, USER)
}

func RegisterRoutes(api api.Service) http.Handler{
	r := mux.NewRouter().PathPrefix(utils.AppParams.ApiBase).Subrouter()
	r.Use(loggingMiddleware)
	for path, config := range routes(api) {
		r.Handle(path, chain.New(config.Middlewares...).Then(config.Handler)).Methods(config.Methods...)
	}
	return r
}
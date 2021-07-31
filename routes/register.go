package routes

import (
	"github.com/anisbhsl/auth-server/middlewares"
	chain "github.com/justinas/alice"
	"github.com/anisbhsl/auth-server/utils"
	"github.com/gorilla/mux"
	"net/http"
)

var(
	loggingMiddleware=middlewares.LoggingMiddleware
)

func RegisterRoutes() http.Handler{
	r := mux.NewRouter().PathPrefix(utils.AppParams.ApiBase).Subrouter()

	//register globally applicable middlewares here
	r.Use(loggingMiddleware)

	for path, config := range routes() {
		r.Handle(path, chain.New(config.Middlewares...).Then(config.Handler)).Methods(config.Methods...)
	}

	return r
}

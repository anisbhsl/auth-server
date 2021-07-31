package routes

import (
	"github.com/anisbhsl/auth-server/handlers"
	chain "github.com/justinas/alice"
	"net/http"
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

type RouteConfig struct {
	Handler     http.Handler
	Methods     []string
	Middlewares []chain.Constructor
}

var routes = func() map[string]RouteConfig {
	GENERAL := map[string]RouteConfig{
		"/": {
			Handler:     handlers.Index(),
			Methods:     []string{httpMethods.GET},
			Middlewares: nil,
		},
	}

	return func(routeMaps ...map[string]RouteConfig) map[string]RouteConfig {
		routeDefs := routeMaps[0]
		for i := 1; i < len(routeMaps); i++ {
			for path, config := range routeMaps[i] {
				routeDefs[path] = config
			}
		}
		return routeDefs
	}(GENERAL)
}

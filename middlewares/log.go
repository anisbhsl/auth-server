package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		log.Println(fmt.Sprintf("Path: %s | Method: %s",r.URL.Path,r.Method))
		next.ServeHTTP(w,r)
	})
}


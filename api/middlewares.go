package api

import (
	"context"
	"fmt"
	"github.com/anisbhsl/auth-server/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func (s *service) AuthMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		authHeader:=strings.Split(r.Header.Get("Authorization"),"Bearer ")
		if len(authHeader)!=2{
			SendUnauthorizedResponse(w)
			return
		}

		if claims,err:=s.AuthService.VerifyToken(authHeader[1]);err!=nil{
			logger.Error(fmt.Sprintf("%v",err))
			SendUnauthorizedResponse(w)
			return
		}else{
			r=r.WithContext(context.WithValue(r.Context(),"email",claims.Email))
		}
		next.ServeHTTP(w,r)
	})
}

func (s *service) LoggingMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		traceID:=uuid.New().String()
		logger.Info("Request Received",zap.String("Path",r.URL.Path),zap.String("Method",r.Method),zap.String("traceID",traceID))
		next.ServeHTTP(w,r.WithContext(context.WithValue(r.Context(),"traceID",traceID)))
	})
}


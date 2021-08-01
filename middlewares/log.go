package middlewares

import (
	"context"
	"github.com/anisbhsl/auth-server/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		traceID:=uuid.New().String()
		logger.Info("Request Received",zap.String("Path",r.URL.Path),zap.String("Method",r.Method),zap.String("traceID",traceID))
		next.ServeHTTP(w,r.WithContext(context.WithValue(r.Context(),"traceID",traceID)))
	})
}


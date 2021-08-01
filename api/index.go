package api

import (
	"encoding/json"
	"github.com/anisbhsl/auth-server/logger"
	"net/http"
)

func (s service) Index() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		logger.Info("Index Page",logger.TraceRequestWithContext(r.Context()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success":true,
			"message":"Welcome to Auth Server API",
		})
	}
}



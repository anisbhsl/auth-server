package api

import (
	"encoding/json"
	"net/http"
)

var (
	CodeInvalidLoginCredentials = "invalid_login_credentials"
	CodeInvalidRegistrationData = "invalid_registration_data"
	CodeInvalidRefreshTokenData = "invalid_refresh_token_data"
	CodeErrorTokenGeneration    = "error_token_generation"
	CodeInvalidToken            = "invalid_token"
	CodeEmailAlreadyExists      = "email_already_exists"
)

func SendErrorResponse(w http.ResponseWriter, errorMessage interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": false,
		"error":   errorMessage,
	})
}

func SendSuccessResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    data,
	})
}

func SendUnauthorizedResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
}

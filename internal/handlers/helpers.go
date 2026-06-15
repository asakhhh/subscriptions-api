package handlers

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message" example:"invalid request body"`
}

func respondJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, ErrorResponse{Message: message})
}

package utils

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func GenerateResponse(data any, status int, message string) response {
	return response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func NewResponse(w http.ResponseWriter, data response) error {
	w.WriteHeader(data.Status)
	return json.NewEncoder(w).Encode(data)
}

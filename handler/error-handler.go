package handler

import (
	"encoding/json"
	"mymodule/utils"
	"net/http"
)

type defaultErrorHandler struct {
}

func NewDefaultErrorHandler() *defaultErrorHandler {
	return &defaultErrorHandler{}
}

func (h *defaultErrorHandler) MethodNotAllowedHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := utils.GenerateResponse(nil, 405, "Dont get mess with the method")

		w.WriteHeader(405)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func (h *defaultErrorHandler) NotFoundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := utils.GenerateResponse(nil, 404, "What are u looking for")

		w.WriteHeader(404)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

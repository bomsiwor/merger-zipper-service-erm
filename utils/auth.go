package utils

import "net/http"

func CheckToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check token from header
		token := r.Header.Get("Authorization")

		// Throw error if token is empty
		if token == "" {
			w.Header().Set("Content-Type", "application/json")
			NewResponse(w, GenerateResponse(nil, 401, "Unauthorized"))
			return
		}

		// Pass the middleware
		next.ServeHTTP(w, r)
	})
}

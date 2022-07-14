package middleware

import "net/http"

func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("x-authentication-token")

		if token == "" {
			http.Error(w, "A valid Authentication header is required to make this request", http.StatusUnauthorized)
		}

		next.ServeHTTP(w, r)
	})
}

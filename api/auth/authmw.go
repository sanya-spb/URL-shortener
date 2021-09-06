package auth

import "net/http"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if user, pass, ok := r.BasicAuth(); !ok || !(user == "admin" && pass == "admin") {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		},
	)
}

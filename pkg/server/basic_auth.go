package server

import (
	"crypto/subtle"
	"net/http"
)

func BasicAuth(username, password string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, pass, ok := r.BasicAuth()
			if !ok {
				WriteError(w, nil, http.StatusForbidden, "forbidden", "access denied")
				return
			}

			if subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
				WriteError(w, nil, http.StatusForbidden, "forbidden", "access denied")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

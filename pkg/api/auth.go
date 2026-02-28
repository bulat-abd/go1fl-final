package api

import (
	"net/http"

	"github.com/bulat-abd/go1fl-final/pkg/token"
)

func auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var jwt string
		cookie, err := r.Cookie("token")
		if err == nil {
			jwt = cookie.Value
		}
		var valid bool
		valid = token.ValidateToken(jwt)
		if !valid {
			http.Error(w, "Authentification required", http.StatusUnauthorized)
			return
		}
		next(w, r)
	})
}

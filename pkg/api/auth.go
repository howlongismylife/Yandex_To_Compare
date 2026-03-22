package api

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"os"
)

func auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		pass := os.Getenv("TODO_PASSWORD")

		// если пароль не задан — пропускаем всё
		if pass == "" {
			next(w, r)
			return
		}

		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Authentification required", http.StatusUnauthorized)
			return
		}

		hash := sha256.Sum256([]byte(pass))
		validToken := hex.EncodeToString(hash[:])

		if cookie.Value != validToken {
			http.Error(w, "Authentification required", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

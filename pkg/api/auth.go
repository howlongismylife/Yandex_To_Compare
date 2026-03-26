package api

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
)

func auth(pass string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// если пароль не задан — пропускаем всё
		if pass == "" {
			next(w, r)
			return
		}

		cookie, err := r.Cookie("token")
		if err != nil {
			writeJSON(w, http.StatusUnauthorized, map[string]any{
				"error": "authentication required",
			})
			return
		}

		hash := sha256.Sum256([]byte(pass))
		validToken := hex.EncodeToString(hash[:])

		if cookie.Value != validToken {
			writeJSON(w, http.StatusUnauthorized, map[string]any{
				"error": "authentication required",
			})
			return
		}

		next(w, r)
	}
}

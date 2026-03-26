package api

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
)

type signinReq struct {
	Password string `json:"password"`
}

func signinHandler(pass string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req signinReq

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{
				"error": err.Error(),
			})
			return
		}

		if pass == "" {
			writeJSON(w, http.StatusInternalServerError, map[string]any{
				"error": "password is not configured",
			})
			return
		}

		if req.Password != pass {
			writeJSON(w, http.StatusUnauthorized, map[string]any{
				"error": "invalid password",
			})
			return
		}

		hash := sha256.Sum256([]byte(pass))
		token := hex.EncodeToString(hash[:])

		writeJSON(w, http.StatusOK, map[string]any{
			"token": token,
		})
	}
}

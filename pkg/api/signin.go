package api

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"os"
)

type signinReq struct {
	Password string `json:"password"`
}

func signinHandler(w http.ResponseWriter, r *http.Request) {
	var req signinReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, map[string]any{"error": err.Error()})
		return
	}

	pass := os.Getenv("TODO_PASSWORD")
	if pass == "" {
		writeJSON(w, map[string]any{"error": "Пароль не задан"})
		return
	}

	if req.Password != pass {
		writeJSON(w, map[string]any{"error": "Неверный пароль"})
		return
	}

	hash := sha256.Sum256([]byte(pass))
	token := hex.EncodeToString(hash[:])

	writeJSON(w, map[string]any{"token": token})
}

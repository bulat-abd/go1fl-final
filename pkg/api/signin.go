package api

import (
	"encoding/json"
	"net/http"

	"github.com/bulat-abd/go1fl-final/internal/config"
	"github.com/bulat-abd/go1fl-final/pkg/token"
)

func signinHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var message map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&message)
	defer r.Body.Close()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	password := message["password"].(string)
	// Timing attack vuln ahead!
	if password == config.GetPassword() {
		token, err := token.CreateToken()
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{"token": token})
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Неправильный пароль"})
		return
	}
}

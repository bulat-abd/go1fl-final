package api

import (
	"encoding/json"
	"net/http"

	"github.com/bulat-abd/go1fl-final/internal/config"
	"github.com/bulat-abd/go1fl-final/pkg/token"
)

func signinHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		JsonError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var message map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&message)
	defer r.Body.Close()
	if err != nil {
		JsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	password := message["password"].(string)
	// Timing attack vuln ahead!
	if password == config.GetPassword() {
		token, err := token.CreateToken()
		if err != nil {
			JsonError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		JsonResponse(w, map[string]interface{}{"token": token})
		return
	} else {
		JsonError(w, "Неправильный пароль", http.StatusOK)
		return
	}
}

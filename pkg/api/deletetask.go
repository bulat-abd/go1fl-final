package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bulat-abd/go1fl-final/pkg/db"
)

func deleteTaskHandler(database *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			JsonError(w, "No ID presented", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			JsonError(w, "No ID presented", http.StatusBadRequest)
			return
		}
		ts := db.NewTaskStore(database)
		err = ts.Delete(int64(id))
		if err != nil {
			JsonError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{})
	}
}

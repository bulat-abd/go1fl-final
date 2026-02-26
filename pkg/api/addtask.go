package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/bulat-abd/go1fl-final/pkg/db"
)

func addTaskHandler(database *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var task db.Task
		err := json.NewDecoder(r.Body).Decode(&task)
		defer r.Body.Close()
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		if task.Title == "" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "Bad title",
			})
			return
		}
		err = task.CheckDate()
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		// add task to db
		ts := db.NewTaskStore(database)
		id, err := ts.Add(task)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		// return id in JSON
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"id": id,
		})
	}
}

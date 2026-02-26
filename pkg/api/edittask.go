package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/bulat-abd/go1fl-final/pkg/db"
)

func editTaskHandler(database *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var task db.Task
		err := json.NewDecoder(r.Body).Decode(&task)
		defer r.Body.Close()
		if err != nil {
			JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}
		if task.Title == "" {
			JsonError(w, "Bad title", http.StatusBadRequest)
			return
		}
		err = task.CheckDate()
		if err != nil {
			JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}
		// add task to db
		ts := db.NewTaskStore(database)
		err = ts.Update(task.ID, task.Date, task.Title, task.Comment, task.Repeat)
		if err != nil {
			JsonError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// return id in JSON
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{})
	}
}

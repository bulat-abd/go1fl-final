package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/bulat-abd/go1fl-final/pkg/db"
)

/*
type TasksResp struct {
    Tasks []*db.Task `json:"tasks"`
}
*/

func listTasksHandler(database *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ts := db.NewTaskStore(database)
		tasks, err := ts.Upcoming(50)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		if tasks == nil {
			tasks = make([]db.Task, 0)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"tasks": tasks,
		})
	}
}

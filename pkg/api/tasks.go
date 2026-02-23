package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/bulat-abd/go1fl-final/pkg/db"
)

/*
type TasksResp struct {
    Tasks []*db.Task `json:"tasks"`
}
*/

func listTasksHandler(database *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var tasks []db.Task
		var err error
		ts := db.NewTaskStore(database)
		search := r.URL.Query().Get("search")
		if search == "" {
			tasks, err = ts.Upcoming(50)
		} else {
			t, err := time.Parse("02.01.2006", search)
			if isText := err != nil; isText {
				tasks, err = ts.SearchByText(search, 50)
			} else {
				tasks, err = ts.SearchByDate(t.Format("20060102"), 50)
			}
		}
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

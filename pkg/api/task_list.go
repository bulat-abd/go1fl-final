package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/bulat-abd/go1fl-final/internal/config"
	"github.com/bulat-abd/go1fl-final/pkg/db"
)

var MaxTasks = int64(config.GetMaxTasks())

func listTasksHandler(database *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var tasks []db.Task
		var err error
		ts := db.NewTaskStore(database)
		search := r.URL.Query().Get("search")
		t, err := time.Parse("02.01.2006", search)
		switch {
		case search == "":
			tasks, err = ts.Upcoming(MaxTasks)
		case err != nil:
			tasks, err = ts.SearchByText(search, MaxTasks)
		default:
			tasks, err = ts.SearchByDate(t.Format("20060102"), MaxTasks)
		}

		if err != nil {
			JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}
		if tasks == nil {
			tasks = make([]db.Task, 0)
		}
		JsonResponse(w, map[string]interface{}{"tasks": tasks})
	}
}

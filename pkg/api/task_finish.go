package api

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/bulat-abd/go1fl-final/pkg/datecalc"
	"github.com/bulat-abd/go1fl-final/pkg/db"
)

func finishTaskHandler(database *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			JsonError(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			JsonError(w, "No ID presented", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			JsonError(w, err.Error(), http.StatusBadRequest)
			return
		}
		ts := db.NewTaskStore(database)
		task, err := ts.Get(int64(id))
		if err != nil {
			JsonError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if task.Repeat == "" {
			err = ts.Delete(task.ID)
			if err != nil {
				JsonError(w, err.Error(), http.StatusInternalServerError)
				return
			}
			JsonResponse(w, map[string]interface{}{})
			return
		}
		now := time.Now()
		nextDate, err := datecalc.NextDate(now, task.Date, task.Repeat)
		err = ts.SetDate(int64(id), nextDate)
		if err != nil {
			JsonError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		JsonResponse(w, map[string]interface{}{})
	}
}

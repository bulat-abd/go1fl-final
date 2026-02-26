package api

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/bulat-abd/go1fl-final/pkg/db"
)

func getTaskHandler(database *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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
		JsonResponse(w, task)

	}
}

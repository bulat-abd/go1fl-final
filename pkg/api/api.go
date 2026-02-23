package api

import (
	"database/sql"
	"net/http"
)

func Init(db *sql.DB) {
	http.HandleFunc("/api/nextdate", nextDateHandler)
	http.HandleFunc("/api/task", addTaskHandler(db))
}

/*
func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// обработка других методов будет добавлена на следующих шагах
	case http.MethodPost:
		addTaskHandler(w, r)
	}
}
*/

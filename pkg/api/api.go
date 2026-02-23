package api

import (
	"database/sql"
	"net/http"
)

func Init(db *sql.DB) {
	http.HandleFunc("/api/nextdate", nextDateHandler)
	http.HandleFunc("/api/task", taskHandler(db))
	http.HandleFunc("/api/tasks", tasksHandler(db))
}

func taskHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// обработка других методов будет добавлена на следующих шагах
		case http.MethodPost:
			addTaskHandler(db)(w, r)
		}
	}
}

func tasksHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// обработка других методов будет добавлена на следующих шагах
		case http.MethodGet:
			listTasksHandler(db)(w, r)
		}
	}
}

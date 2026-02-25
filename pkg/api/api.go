package api

import (
	"database/sql"
	"net/http"
)

func Init(db *sql.DB) {
	http.HandleFunc("/api/nextdate", nextDateHandler)
	http.HandleFunc("/api/task", taskHandler(db))
	http.HandleFunc("/api/task/done", finishTaskHandler(db))
	http.HandleFunc("/api/tasks", tasksHandler(db))
	http.HandleFunc("/api/signin", signinHandler)
}

func taskHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// обработка других методов будет добавлена на следующих шагах
		case http.MethodPost:
			addTaskHandler(db)(w, r)
		case http.MethodGet:
			getTaskHandler(db)(w, r)
		case http.MethodPut:
			editTaskHandler(db)(w, r)
		case http.MethodDelete:
			deleteTaskHandler(db)(w, r)
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

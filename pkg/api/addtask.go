package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/bulat-abd/go1fl-final/pkg/datecalc"
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
		err = checkDate(&task)
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

func checkDate(task *db.Task) error {
	var next string
	now := time.Now()
	if task.Date == "" {
		task.Date = now.Format("20060102")
	}
	t, err := time.Parse("20060102", task.Date)
	if err != nil {
		return err
	}
	if task.Repeat != "" {
		next, err = datecalc.NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}

	}
	if datecalc.AfterNow(now, t) {
		if len(task.Repeat) == 0 {
			// если правила повторения нет, то берём сегодняшнее число
			task.Date = now.Format("20060102")
		} else {
			// в противном случае, берём вычисленную ранее следующую дату
			task.Date = next
		}
	}
	return nil
}

package db

import (
	"time"

	"github.com/bulat-abd/go1fl-final/pkg/datecalc"
)

type Task struct {
	ID      int64  `db:"id" json:"id,string"`
	Date    string `db:"date" json:"date"`
	Title   string `db:"title" json:"title"`
	Comment string `db:"comment" json:"comment"`
	Repeat  string `db:"repeat" json:"repeat"`
}

func (task *Task) CheckDate() error {
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

type TaskService struct {
	store TaskStore
}

func NewTaskService(store TaskStore) TaskService {
	return TaskService{store: store}
}

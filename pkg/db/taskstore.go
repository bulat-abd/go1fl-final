package db

import (
	"database/sql"
)

type TaskStore struct {
	db *sql.DB
}

func NewTaskStore(db *sql.DB) TaskStore {
	return TaskStore{db: db}
}

func (ts TaskStore) Add(t Task) (int64, error) {
	res, err := ts.db.Exec(
		"INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)",
		sql.Named("date", t.Date),
		sql.Named("title", t.Title),
		sql.Named("comment", t.Comment),
		sql.Named("repeat", t.Repeat),
	)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return id, err
}

func (ts TaskStore) Get(id int64) (Task, error) {
	row := ts.db.QueryRow("SELECT id, date, title, comment, repeat FROM scheduler WHERE id = :id", sql.Named("id", id))
	t := Task{}
	err := row.Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat)
	if err != nil {
		return Task{}, err
	}
	return t, nil
}

func (ts TaskStore) Delete(id int64) error {
	_, err := ts.db.Exec(
		"DELETE FROM scheduler WHERE id = :id",
		sql.Named("id", id),
	)
	return err
}

func (ts TaskStore) SetDate(id int64, date string) error {
	_, err := ts.db.Exec(
		"UPDATE scheduler SET date = :date WHERE id = :id",
		sql.Named("date", date),
		sql.Named("id", id))
	return err
}

func (ts TaskStore) Update(id int64, date, title, comment, repeat string) error {
	_, err := ts.db.Exec(
		"UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id",
		sql.Named("date", date),
		sql.Named("title", title),
		sql.Named("comment", comment),
		sql.Named("repeat", repeat),
		sql.Named("id", id))
	return err
}

func (ts TaskStore) Upcoming(count int64) ([]Task, error) {
	rows, err := ts.db.Query("SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date ASC LIMIT :count", sql.Named("count", count))
	if err != nil {
		return nil, err
	}
	tasks := make([]Task, 0, count)
	for rows.Next() {
		var task Task
		// Scan the row data into the struct fields
		if err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (ts TaskStore) SearchByDate(date string, count int64) ([]Task, error) {
	rows, err := ts.db.Query("SELECT id, date, title, comment, repeat FROM scheduler WHERE date = :date LIMIT :count", sql.Named("date", date), sql.Named("count", count))
	if err != nil {
		return nil, err
	}
	tasks := make([]Task, 0, count)
	for rows.Next() {
		var task Task
		// Scan the row data into the struct fields
		if err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (ts TaskStore) SearchByText(text string, count int64) ([]Task, error) {
	pattern := "%" + text + "%"
	rows, err := ts.db.Query("SELECT id, date, title, comment, repeat FROM scheduler WHERE title LIKE :pattern OR comment LIKE :pattern ORDER BY date ASC LIMIT :count", sql.Named("pattern", pattern), sql.Named("count", count))
	if err != nil {
		return nil, err
	}
	tasks := make([]Task, 0, count)
	for rows.Next() {
		var task Task
		// Scan the row data into the struct fields
		if err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

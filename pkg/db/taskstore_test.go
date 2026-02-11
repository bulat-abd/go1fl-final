package db

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// getTestParcel возвращает тестовую посылку
func getTestTask() Task {
	return Task{
		Date:    "20060102",
		Title:   "test title",
		Comment: "test comment",
		Repeat:  "w 1,4,5",
	}
}

func TestAddGetDelete(t *testing.T) {
	db, err := sql.Open("sqlite", "file:test.db?mode=memory&cache=shared")
	require.NoError(t, err)
	defer db.Close()
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS "scheduler"
	(
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
		date CHAR(8) NOT NULL DEFAULT "",
    	title varchar(128) NOT NULL,
        comment text NOT NULL,
        repeat varchar(128) NOT NULL
	);`
	_, err = db.Exec(createTableSQL)
	require.NoError(t, err)
	store := NewTaskStore(db)
	task := getTestTask()

	id, err := store.Add(task)
	require.NoError(t, err)
	require.NotEmpty(t, id)

	storedTask, err := store.Get(id)
	require.NoError(t, err)
	assert.Equal(t, storedTask.ID, id)
	assert.Equal(t, storedTask.Date, task.Date)
	assert.Equal(t, storedTask.Title, task.Title)
	assert.Equal(t, storedTask.Comment, task.Comment)
	assert.Equal(t, storedTask.Repeat, task.Repeat)

	err = store.Delete(id)
	require.NoError(t, err)
	_, err = store.Get(id)
	require.Error(t, err)
}

func TestSetDate(t *testing.T) {
	db, err := sql.Open("sqlite", "file:test.db?mode=memory&cache=shared")
	require.NoError(t, err)
	defer db.Close()
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS "scheduler"
	(
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
		date CHAR(8) NOT NULL DEFAULT "",
    	title varchar(128) NOT NULL,
        comment text NOT NULL,
        repeat varchar(128) NOT NULL
	);`
	_, err = db.Exec(createTableSQL)
	require.NoError(t, err)
	store := NewTaskStore(db)
	task := getTestTask()

	id, err := store.Add(task)
	require.NoError(t, err)
	require.NotEmpty(t, id)

	err = store.SetDate(id, "13372048")
	storedTask, err := store.Get(id)
	require.NoError(t, err)
	assert.Equal(t, storedTask.Date, "13372048")
}

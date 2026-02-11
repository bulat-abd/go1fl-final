package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

var FilePath = "scheduler.db"

func Init(dbFile string) (*sql.DB, error) {
	_, err := os.Stat(dbFile)
	var install bool
	if err != nil {
		install = true
	}

	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
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
	if install {
		_, err := db.Exec(createTableSQL)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return db, nil
}

package main

import (
	"github.com/bulat-abd/go1fl-final/pkg/db"
	"github.com/bulat-abd/go1fl-final/pkg/server"
)

func main() {
	database, _ := db.Init()
	defer database.Close()
	server.Run(database)
}

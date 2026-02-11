package main

import (
	"github.com/bulat-abd/go1fl-final/pkg/db"
	"github.com/bulat-abd/go1fl-final/pkg/server"
)

func main() {
	_, _ = db.Init(db.FilePath)
	server.Run()
}

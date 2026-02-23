package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/bulat-abd/go1fl-final/internal/config"
	"github.com/bulat-abd/go1fl-final/pkg/api"
)

func Run(db *sql.DB) {
	http.Handle("/", http.FileServer(http.Dir(config.WebDir)))
	api.Init(db)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.GetPort()), nil)
	if err != nil {
		panic(err)
	}
}

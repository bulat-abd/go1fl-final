package server

import (
	"fmt"
	"net/http"

	"github.com/bulat-abd/go1fl-final/internal/config"
)

func Run() {
	http.Handle("/", http.FileServer(http.Dir(config.WebDir)))
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.GetPort()), nil)
	if err != nil {
		panic(err)
	}
}

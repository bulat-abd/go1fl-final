package main

import (
	"fmt"
	"net/http"
)

var port = 7540
var webDir = "./web"

func main() {
	http.Handle("/", http.FileServer(http.Dir(webDir)))
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var port = 7540
var webDir = "./web"

func getPort(defaultPort int) int {
	port := defaultPort
	envPort := os.Getenv("TODO_PORT")
	if len(envPort) > 0 {
		if parsedPort, err := strconv.ParseInt(envPort, 10, 32); err == nil {
			port = int(parsedPort)
		}
	}
	return port
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(webDir)))
	err := http.ListenAndServe(fmt.Sprintf(":%d", getPort(port)), nil)
	if err != nil {
		panic(err)
	}
}

package config

import (
	"os"
	"strconv"
)

var defaultPort = 7540
var WebDir = "./web"

func GetPort() int {
	port := defaultPort
	envPort := os.Getenv("TODO_PORT")
	if len(envPort) > 0 {
		if parsedPort, err := strconv.ParseInt(envPort, 10, 32); err == nil {
			port = int(parsedPort)
		}
	}
	return port
}

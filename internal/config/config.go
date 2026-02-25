package config

import (
	"os"
	"strconv"
)

var defaultPort = 7540
var WebDir = "./web"

var defaultDBFilePath = "scheduler.db"
var defaultMaxTasks = 50
var defaultPassword = "12345"

var Secret = []byte("secret")

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

func GetDBFilePath() string {
	path := defaultDBFilePath
	envPath := os.Getenv("TODO_DBFILE")
	if len(envPath) > 0 {
		path = envPath
	}
	return path
}

func GetMaxTasks() int {
	maxTasks := defaultMaxTasks
	envMaxTasks := os.Getenv("TODO_MAXTASKS")
	if len(envMaxTasks) > 0 {
		if parsedMaxTasks, err := strconv.ParseInt(envMaxTasks, 10, 32); err == nil {
			maxTasks = int(parsedMaxTasks)
		}
	}
	return maxTasks
}

func GetPassword() string {
	password := defaultPassword
	envPassword := os.Getenv("TODO_PASSWORD")
	if len(envPassword) > 0 {
		password = envPassword
	}
	return password
}

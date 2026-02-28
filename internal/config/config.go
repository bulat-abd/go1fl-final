package config

import (
	"os"
	"strconv"
)

const (
	defaultDBFilePath  = "scheduler.db"
	defaultMaxTasks    = 50
	defaultPassword    = "12345"
	defaultPort        = 7540
	defaultTokenSecret = "secret" // gets converted to []byte, should maybe used base64/hex encoding here

	WebDir = "./web"
)

func getIntParameter(envVar string, defaultValue int) int {
	result := defaultValue
	envValueStr := os.Getenv(envVar)
	if len(envValueStr) > 0 {
		if parsedValue, err := strconv.ParseInt(envValueStr, 10, 32); err == nil {
			result = int(parsedValue)
		}

	}
	return result
}

func GetPort() int {
	return getIntParameter("TODO_PORT", defaultPort)
}

func GetMaxTasks() int {
	return getIntParameter("TODO_MAXTASKS", defaultMaxTasks)
}

func getStringParameter(envVar string, defaultValue string) string {
	result := defaultValue
	envValueStr := os.Getenv(envVar)
	if len(envValueStr) > 0 {
		result = envValueStr
	}
	return result
}

func GetDBFilePath() string {
	return getStringParameter("TODO_DBFILE", defaultDBFilePath)
}

func GetPassword() string {
	return getStringParameter("TODO_PASSWORD", defaultPassword)
}

func GetTokenSecret() []byte {
	return []byte(getStringParameter("TODO_TOKENSECRET", defaultTokenSecret))
}

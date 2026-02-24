package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigWithoutEnv(t *testing.T) {
	assert.Equal(t, defaultPort, GetPort(), `получен неправильный номер порта`)
}

func TestConfigWithEnv(t *testing.T) {
	portValue := 31337
	envValue := fmt.Sprintf("%d", portValue)
	t.Setenv("TODO_PORT", envValue)
	assert.Equal(t, portValue, GetPort(), `получен неправильный номер порта`)
}

func TestDBFilePathWithoutEnv(t *testing.T) {
	assert.Equal(t, defaultDBFilePath, GetDBFilePath(), `получен неправильный путь до файла БД`)
}

func TestDBFilePathWithEnv(t *testing.T) {
	envValue := "/some/path"
	t.Setenv("TODO_DBFILE", envValue)
	assert.Equal(t, envValue, GetDBFilePath(), `получен неправильный путь до файла БД`)
}

func TestMaxTasksWithoutEnv(t *testing.T) {
	assert.Equal(t, defaultMaxTasks, GetMaxTasks(), `получено неправильное максимальное количество задач`)
}

func TestMaxTasksWithEnv(t *testing.T) {
	maxTasks := 100
	envValue := fmt.Sprintf("%d", maxTasks)

	t.Setenv("TODO_MAXTASKS", envValue)
	assert.Equal(t, maxTasks, GetMaxTasks(), `получено неправильное максимальное количество задач`)
}

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

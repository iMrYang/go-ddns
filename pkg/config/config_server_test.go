package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	serverConfig := NewServer()
	assert.NotEmpty(t, serverConfig.Host, "Host is '%v'", serverConfig.Host)
	assert.Less(t, serverConfig.Port, 65536, "Port is '%v', must less than 65536", serverConfig.Port)
	assert.Less(t, 0, serverConfig.Port, "Port is '%v', must more than 0", serverConfig.Port)
}

func TestServerFlagSet(t *testing.T) {
	serverOptions := []string{
		"--server.host", "0.0.0.0",
		"--server.port", "80",
	}
	serverConfig := NewServer()
	assert.Nil(t, serverConfig.FlagSet().Parse(serverOptions), "parse command line")

	assert.Equal(t, "0.0.0.0", serverConfig.Host, "expect debug but Host is '%v'", serverConfig.Host)
	assert.Equal(t, 80, serverConfig.Port, "expect stdout but Port is '%v'", serverConfig.Port)
}

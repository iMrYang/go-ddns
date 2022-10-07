package config

import (
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

type config struct {
	// Log is the log config.
	Log *Log
	// Server is the server config.
	Server *Server
}

const cfgContent = `
log:
  level: debug
  output: stdout
server:
  host: 0.0.0.0
  port: 80
`

func newConfig() *config {
	return &config{
		Log:    NewLog(),
		Server: NewServer(),
	}
}

func TestConfig(t *testing.T) {
	// read config
	v := viper.New()
	v.SetConfigType("yaml")
	v.ReadConfig(strings.NewReader(cfgContent))

	// parse config
	cfg := newConfig()
	v.Unmarshal(&cfg)

	// check log config
	assert.Equal(t, "debug", cfg.Log.Level)
	assert.Equal(t, "stdout", cfg.Log.Output)

	// check server config
	assert.Equal(t, "0.0.0.0", cfg.Server.Host)
	assert.Equal(t, 80, cfg.Server.Port)
}

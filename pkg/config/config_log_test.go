package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLog(t *testing.T) {
	logConfig := NewLog()
	assert.NotEmpty(t, logConfig.Level, "Level is '%v'", logConfig.Level)
	assert.NotEmpty(t, logConfig.Output, "Output is '%v'", logConfig.Output)
}

func TestLogFlagSet(t *testing.T) {
	logOptions := []string{
		"--log.level", "debug",
		"--log.output", "stdout",
	}
	logConfig := NewLog()
	assert.Nil(t, logConfig.FlagSet().Parse(logOptions), "parse command line")

	assert.Equal(t, "debug", logConfig.Level, "expect debug but Level is '%v'", logConfig.Level)
	assert.Equal(t, "stdout", logConfig.Output, "expect stdout but Output is '%v'", logConfig.Output)
}

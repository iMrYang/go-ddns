package config

import "flag"

// Log is the Config interface.
var _ Config = (*Log)(nil)

// Log is the configuration for the log.
type Log struct {
	// Level is the log level.
	Level string `mapstructure:"level" validate:"oneof=debug info warn error"`
	// Output is the log output.
	Output string `mapstructure:"output" validate:""`
}

func NewLog() *Log {
	return &Log{
		Level:  "error",
		Output: "stderr",
	}
}

func (l *Log) FlagSet() *flag.FlagSet {
	fs := flag.NewFlagSet("log", flag.ExitOnError)
	fs.StringVar(&l.Level, "log.level", l.Level, "log level")
	fs.StringVar(&l.Output, "log.output", l.Output, "log output")
	return fs
}

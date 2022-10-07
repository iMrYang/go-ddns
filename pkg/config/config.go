package config

import "flag"

// Config is the interface.
type Config interface {
	FlagSet() *flag.FlagSet
}

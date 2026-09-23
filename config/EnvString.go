package config

import (
	"github.com/kduong-dev/goutil/fatal"
	"os"
)

func EnvString(key string, dflt string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return dflt
}

func EnvStringOrFatal(key string) string {
	value := os.Getenv(key)
	if value == "" {
		fatal.LogErrorf("The following environment variable must be set: %s", key)
	}
	return value
}

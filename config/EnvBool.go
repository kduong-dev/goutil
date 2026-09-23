package config

import (
	"os"
	"strconv"

	"github.com/kduong-dev/goutil/fatal"
)

func EnvBool(key string, dflt bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return dflt
	}
	parsed, err := strconv.ParseBool(value)
	fatal.OnErrorf(err, "environment variable %s must be a bool", key)
	return parsed
}

func EnvBoolOrFatal(key string) bool {
	value := os.Getenv(key)
	if value == "" {
		fatal.LogErrorf("The following environment variable must be set: %s", key)
	}
	parsed, err := strconv.ParseBool(value)
	fatal.OnErrorf(err, "environment variable %s must be a bool", key)
	return parsed
}

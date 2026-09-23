package config

import (
	"os"
	"time"

	"github.com/kduong-dev/goutil/fatal"
)

func EnvDuration(key string, dflt time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return dflt
	}
	parsed, err := time.ParseDuration(value)
	fatal.OnErrorf(err, "environment variable %s must be a valid duration", key)
	return parsed
}

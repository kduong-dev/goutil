package config

import (
	"net/url"
	"os"

	"goutil/fatal"
)

func EnvURLOrFatal(key string) *url.URL {
	value := os.Getenv(key)
	if value == "" {
		fatal.LogErrorf("The following environment variable must be set: %s", key)
	}
	parsed, err := url.Parse(value)
	fatal.OnErrorf(err, "environment variable %s must be a valid URL", key)
	return parsed
}

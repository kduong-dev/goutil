package config

import (
	"goutil/fatal"
	"os"
	"strconv"
)

type float interface {
	~float32 | ~float64
}

func envFloat[T float](key string, dflt T, bitSize int, typeName string) T {
	value := os.Getenv(key)
	if value == "" {
		return dflt
	}
	parsed, err := strconv.ParseFloat(value, bitSize)
	fatal.OnErrorf(err, "environment variable %s must be type %s", key, typeName)
	return T(parsed)
}

func envFloatOrFatal[T float](key string, bitSize int, typeName string) T {
	value := os.Getenv(key)
	if value == "" {
		fatal.LogErrorf("The following environment variable must be set: %s", key)
	}
	parsed, err := strconv.ParseFloat(value, bitSize)
	fatal.OnErrorf(err, "environment variable %s must be type %s", key, typeName)
	return T(parsed)
}

func EnvFloat32(key string, dflt float32) float32 {
	return envFloat(key, dflt, 32, "float32")
}

func EnvFloat32OrFatal(key string) float32 {
	return envFloatOrFatal[float32](key, 32, "float32")
}

func EnvFloat64(key string, dflt float64) float64 {
	return envFloat(key, dflt, 64, "float64")
}

func EnvFloat64OrFatal(key string) float64 {
	return envFloatOrFatal[float64](key, 64, "float64")
}

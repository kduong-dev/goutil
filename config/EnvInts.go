package config

import (
	"github.com/kduong-dev/goutil/fatal"
	"os"
	"strconv"
)

type signedInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func envInt[T signedInt](key string, dflt T, bitSize int, typeName string) T {
	value := os.Getenv(key)
	if value == "" {
		return dflt
	}
	parsed, err := strconv.ParseInt(value, 10, bitSize)
	fatal.OnErrorf(err, "environment variable %s must be type %s", key, typeName)
	return T(parsed)
}

func envIntOrFatal[T signedInt](key string, bitSize int, typeName string) T {
	value := os.Getenv(key)
	if value == "" {
		fatal.LogErrorf("The following environment variable must be set: %s", key)
	}
	parsed, err := strconv.ParseInt(value, 10, bitSize)
	fatal.OnErrorf(err, "environment variable %s must be type %s", key, typeName)
	return T(parsed)
}

func EnvInt(key string, dflt int) int {
	return envInt(key, dflt, strconv.IntSize, "int")
}

func EnvIntOrFatal(key string) int {
	return envIntOrFatal[int](key, strconv.IntSize, "int")
}

func EnvInt8(key string, dflt int8) int8 {
	return envInt(key, dflt, 8, "int8")
}

func EnvInt8OrFatal(key string) int8 {
	return envIntOrFatal[int8](key, 8, "int8")
}

func EnvInt16(key string, dflt int16) int16 {
	return envInt(key, dflt, 16, "int16")
}

func EnvInt16OrFatal(key string) int16 {
	return envIntOrFatal[int16](key, 16, "int16")
}

func EnvInt32(key string, dflt int32) int32 {
	return envInt(key, dflt, 32, "int32")
}

func EnvInt32OrFatal(key string) int32 {
	return envIntOrFatal[int32](key, 32, "int32")
}

func EnvInt64(key string, dflt int64) int64 {
	return envInt(key, dflt, 64, "int64")
}

func EnvInt64OrFatal(key string) int64 {
	return envIntOrFatal[int64](key, 64, "int64")
}

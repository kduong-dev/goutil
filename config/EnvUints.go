package config

import (
	"github.com/kduong-dev/goutil/fatal"
	"os"
	"strconv"
)

type unsignedInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func envUint[T unsignedInt](key string, dflt T, bitSize int, typeName string) T {
	value := os.Getenv(key)
	if value == "" {
		return dflt
	}
	parsed, err := strconv.ParseUint(value, 10, bitSize)
	fatal.OnErrorf(err, "environment variable %s must be type %s", key, typeName)
	return T(parsed)
}

func envUintOrFatal[T unsignedInt](key string, bitSize int, typeName string) T {
	value := os.Getenv(key)
	if value == "" {
		fatal.LogErrorf("The following environment variable must be set: %s", key)
	}
	parsed, err := strconv.ParseUint(value, 10, bitSize)
	fatal.OnErrorf(err, "environment variable %s must be type %s", key, typeName)
	return T(parsed)
}

func EnvUint(key string, dflt uint) uint {
	return envUint(key, dflt, strconv.IntSize, "uint")
}

func EnvUintOrFatal(key string) uint {
	return envUintOrFatal[uint](key, strconv.IntSize, "uint")
}

func EnvUint32(key string, dflt uint32) uint32 {
	return envUint(key, dflt, 32, "uint32")
}

func EnvUint32OrFatal(key string) uint32 {
	return envUintOrFatal[uint32](key, 32, "uint32")
}

func EnvUint64(key string, dflt uint64) uint64 {
	return envUint(key, dflt, 64, "uint64")
}

func EnvUint64OrFatal(key string) uint64 {
	return envUintOrFatal[uint64](key, 64, "uint64")
}

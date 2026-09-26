package config

import "net/url"

// EnvURLFromPartsOrFatal builds a URL from the prefix_SCHEME, prefix_HOST and
// prefix_PORT environment variables. prefix_PORT may be empty.
func EnvURLFromPartsOrFatal(prefix string) *url.URL {
	host := EnvStringOrFatal(prefix + "_HOST")
	if port := EnvString(prefix+"_PORT", ""); port != "" {
		host += ":" + port
	}
	return &url.URL{Scheme: EnvStringOrFatal(prefix + "_SCHEME"), Host: host}
}

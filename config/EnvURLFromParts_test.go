package config_test

import (
	"testing"

	"github.com/kduong-dev/goutil/config"
	. "github.com/smartystreets/goconvey/convey"
)

func TestEnvURLFromPartsOrFatal(t *testing.T) {
	Convey("Given scheme, host and port environment variables", t, func() {
		t.Setenv("JOURNAL_SERVICE_SCHEME", "http")
		t.Setenv("JOURNAL_SERVICE_HOST", "journal-service")
		t.Setenv("JOURNAL_SERVICE_PORT", "8080")
		Convey("When the URL is built", func() {
			built := config.EnvURLFromPartsOrFatal("JOURNAL_SERVICE")
			Convey("Then it joins the host and port", func() {
				So(built.String(), ShouldEqual, "http://journal-service:8080")
			})
		})
	})
	Convey("Given scheme and host environment variables without a port", t, func() {
		t.Setenv("ALPACA_API_SCHEME", "https")
		t.Setenv("ALPACA_API_HOST", "api.alpaca.markets")
		t.Setenv("ALPACA_API_PORT", "")
		Convey("When the URL is built", func() {
			built := config.EnvURLFromPartsOrFatal("ALPACA_API")
			Convey("Then it omits the port", func() {
				So(built.String(), ShouldEqual, "https://api.alpaca.markets")
			})
		})
	})
}

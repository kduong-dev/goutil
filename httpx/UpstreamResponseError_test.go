package httpx_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/ansel1/merry"
	"github.com/kduong-dev/goutil/httpx"
	. "github.com/smartystreets/goconvey/convey"
)

func upstreamResponse(statusCode int, body string) *http.Response {
	return &http.Response{StatusCode: statusCode, Body: io.NopCloser(strings.NewReader(body))}
}

func TestUpstreamResponseError(t *testing.T) {
	Convey("Given a successful upstream response", t, func() {
		response := upstreamResponse(http.StatusOK, `{}`)
		Convey("When its error is extracted", func() {
			err := httpx.UpstreamResponseError(response)
			Convey("Then there is no error", func() {
				So(err, ShouldBeNil)
			})
		})
	})
	Convey("Given an upstream response rejecting our credentials", t, func() {
		response := upstreamResponse(http.StatusUnauthorized, `{"message":"invalid api key"}`)
		Convey("When its error is extracted", func() {
			err := httpx.UpstreamResponseError(response)
			Convey("Then it is reported as a bad gateway rather than unauthorized", func() {
				So(merry.HTTPCode(err), ShouldEqual, http.StatusBadGateway)
			})
			Convey("Then the upstream user message is kept", func() {
				So(merry.UserMessage(err), ShouldEqual, "invalid api key")
			})
		})
	})
}

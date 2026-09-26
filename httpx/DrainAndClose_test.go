package httpx_test

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/kduong-dev/goutil/httpx"
	. "github.com/smartystreets/goconvey/convey"
)

type readCloserFixture struct {
	io.Reader
	closed bool
}

func (readCloser *readCloserFixture) Close() error {
	readCloser.closed = true
	return nil
}

type failingReader struct{}

func (failingReader) Read([]byte) (int, error) {
	return 0, errors.New("connection reset")
}

func TestDrainAndClose(t *testing.T) {
	Convey("Given a partially read body", t, func() {
		body := &readCloserFixture{Reader: strings.NewReader("unread response body")}
		Convey("When it is drained and closed", func() {
			err := httpx.DrainAndClose(body)
			Convey("Then the rest of the body is consumed and it is closed", func() {
				So(err, ShouldBeNil)
				remaining, _ := io.ReadAll(body.Reader)
				So(remaining, ShouldBeEmpty)
				So(body.closed, ShouldBeTrue)
			})
		})
	})
	Convey("Given a body that fails to read", t, func() {
		body := &readCloserFixture{Reader: failingReader{}}
		Convey("When it is drained and closed", func() {
			err := httpx.DrainAndClose(body)
			Convey("Then the read error is returned and it is still closed", func() {
				So(err, ShouldNotBeNil)
				So(body.closed, ShouldBeTrue)
			})
		})
	})
}

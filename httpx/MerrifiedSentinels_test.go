package httpx_test

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/ansel1/merry"
	"github.com/kduong-dev/goutil/httpx"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMerrifiedSentinels(t *testing.T) {
	Convey("Given merrified sentinels for a missing book and an invalid cursor", t, func() {
		errBookNotFound := errors.New("book not found")
		errInvalidCursor := errors.New("invalid cursor")
		merrifiedSentinels := httpx.MerrifiedSentinels{
			{Sentinel: errBookNotFound, StatusCode: http.StatusNotFound, UserMessage: "no book with that id"},
			{Sentinel: errInvalidCursor, StatusCode: http.StatusBadRequest, UserMessage: "invalid cursor"},
		}
		Convey("When a wrapped sentinel is merrified", func() {
			err := merrifiedSentinels.Merrify(fmt.Errorf("get book: %w", errBookNotFound))
			Convey("Then it carries the sentinel's status code and user message", func() {
				So(merry.HTTPCode(err), ShouldEqual, http.StatusNotFound)
				So(merry.UserMessage(err), ShouldEqual, "no book with that id")
			})
			Convey("Then it still matches the sentinel", func() {
				So(errors.Is(err, errBookNotFound), ShouldBeTrue)
			})
		})
		Convey("When an unknown error is merrified", func() {
			unknownErr := errors.New("disk full")
			err := merrifiedSentinels.Merrify(unknownErr)
			Convey("Then it is returned unchanged", func() {
				So(err, ShouldEqual, unknownErr)
			})
		})
		Convey("When nil is merrified", func() {
			Convey("Then Merrify returns nil", func() {
				So(merrifiedSentinels.Merrify(nil), ShouldBeNil)
			})
			Convey("Then MerrifyOrFatal returns nil", func() {
				So(merrifiedSentinels.MerrifyOrFatal(nil), ShouldBeNil)
			})
		})
		Convey("When a sentinel is merrified or fatal", func() {
			err := merrifiedSentinels.MerrifyOrFatal(errInvalidCursor)
			Convey("Then it carries the sentinel's status code and user message", func() {
				So(merry.HTTPCode(err), ShouldEqual, http.StatusBadRequest)
				So(merry.UserMessage(err), ShouldEqual, "invalid cursor")
			})
		})
	})
}

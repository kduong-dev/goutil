package httpx_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kduong-dev/goutil/httpx"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHandlerWithCORS(t *testing.T) {
	Convey("Given a handler wrapped with CORS", t, func() {
		nextCalled := false
		handler := httpx.HandlerWithCORS(http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
			nextCalled = true
			responseWriter.WriteHeader(http.StatusOK)
		}))
		recorder := httptest.NewRecorder()
		Convey("When a preflight request is made", func() {
			request := httptest.NewRequest(http.MethodOptions, "/resource", nil)
			request.Header.Set("Origin", "https://example.com")
			request.Header.Set("Access-Control-Request-Method", http.MethodPatch)
			request.Header.Set("Access-Control-Request-Headers", "authorization,content-type")
			handler.ServeHTTP(recorder, request)
			Convey("Then it is answered without calling the next handler", func() {
				So(nextCalled, ShouldBeFalse)
				So(recorder.Code, ShouldEqual, http.StatusNoContent)
				So(recorder.Header().Get("Access-Control-Allow-Origin"), ShouldEqual, "*")
				So(recorder.Header().Get("Access-Control-Allow-Methods"), ShouldEqual, http.MethodPatch)
				So(recorder.Header().Get("Access-Control-Allow-Headers"), ShouldEqual, "authorization,content-type")
			})
		})
		Convey("When a cross-origin request is made", func() {
			request := httptest.NewRequest(http.MethodGet, "/resource", nil)
			request.Header.Set("Origin", "https://example.com")
			handler.ServeHTTP(recorder, request)
			Convey("Then the next handler is called and the CORS headers are set", func() {
				So(nextCalled, ShouldBeTrue)
				So(recorder.Code, ShouldEqual, http.StatusOK)
				So(recorder.Header().Get("Access-Control-Allow-Origin"), ShouldEqual, "*")
				So(recorder.Header().Get("Access-Control-Expose-Headers"), ShouldContainSubstring, "Content-Range")
			})
		})
		Convey("When a same-origin request is made", func() {
			request := httptest.NewRequest(http.MethodGet, "/resource", nil)
			handler.ServeHTTP(recorder, request)
			Convey("Then the next handler is called without CORS headers", func() {
				So(nextCalled, ShouldBeTrue)
				So(recorder.Header().Get("Access-Control-Allow-Origin"), ShouldBeEmpty)
			})
		})
	})
}

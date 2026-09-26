package httpx_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kduong-dev/goutil/httpx"
	. "github.com/smartystreets/goconvey/convey"
)

type failingResponseWriter struct {
	*httptest.ResponseRecorder
}

func (writer failingResponseWriter) Write([]byte) (int, error) {
	return 0, errors.New("client disconnected")
}

func TestSendJSONResponse(t *testing.T) {
	Convey("Given a JSON-encodable body", t, func() {
		body := map[string]string{"name": "report.html"}
		responseRecorder := httptest.NewRecorder()
		Convey("When the JSON response is sent", func() {
			httpx.SendJSONResponse(responseRecorder, http.StatusCreated, body)
			Convey("Then it writes the status code, JSON content type and newline-terminated body", func() {
				So(responseRecorder.Code, ShouldEqual, http.StatusCreated)
				So(responseRecorder.Header().Get("Content-Type"), ShouldEqual, "application/json")
				So(responseRecorder.Body.String(), ShouldEqual, "{\"name\":\"report.html\"}\n")
			})
		})
	})
	Convey("Given a nil body", t, func() {
		responseRecorder := httptest.NewRecorder()
		Convey("When the JSON response is sent", func() {
			httpx.SendJSONResponse(responseRecorder, http.StatusNoContent, nil)
			Convey("Then it writes only the status code", func() {
				So(responseRecorder.Code, ShouldEqual, http.StatusNoContent)
				So(responseRecorder.Body.Len(), ShouldEqual, 0)
			})
		})
	})
	Convey("Given a body that cannot be marshalled", t, func() {
		body := map[string]any{"channel": make(chan int)}
		responseRecorder := httptest.NewRecorder()
		Convey("When the JSON response is sent", func() {
			send := func() { httpx.SendJSONResponse(responseRecorder, http.StatusOK, body) }
			Convey("Then it panics before writing the status code", func() {
				So(send, ShouldPanic)
				So(responseRecorder.Flushed, ShouldBeFalse)
				So(responseRecorder.Body.Len(), ShouldEqual, 0)
			})
		})
	})
	Convey("Given a client that has disconnected", t, func() {
		responseWriter := failingResponseWriter{ResponseRecorder: httptest.NewRecorder()}
		Convey("When the JSON response is sent", func() {
			send := func() {
				httpx.SendJSONResponse(responseWriter, http.StatusOK, map[string]string{"name": "report.html"})
			}
			Convey("Then the write error is ignored", func() {
				So(send, ShouldNotPanic)
			})
		})
	})
}

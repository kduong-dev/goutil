package fatal_test

import (
	"testing"

	"github.com/kduong-dev/goutil/fatal"
	. "github.com/smartystreets/goconvey/convey"
)

type marshalFixture struct {
	Name string `json:"name"`
}

func TestUnlessMarshal(t *testing.T) {
	Convey("Given a JSON-serialisable value", t, func() {
		value := marshalFixture{Name: "report"}
		Convey("When it is marshalled", func() {
			data := fatal.UnlessMarshal(value)
			Convey("Then it returns the JSON encoding", func() {
				So(string(data), ShouldEqual, `{"name":"report"}`)
			})
			Convey("Then unmarshalling it round-trips the value", func() {
				var decoded marshalFixture
				fatal.UnlessUnmarshal(data, &decoded)
				So(decoded, ShouldResemble, value)
			})
		})
	})
}

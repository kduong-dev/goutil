package fatal_test

import (
	"context"
	"errors"
	"testing"

	"github.com/kduong-dev/goutil/fatal"
	. "github.com/smartystreets/goconvey/convey"
)

func TestOnErrorUnlessDone(t *testing.T) {
	Convey("Given an active context", t, func() {
		ctx := context.Background()
		Convey("When there is no error", func() {
			Convey("Then it returns without exiting", func() {
				So(func() { fatal.OnErrorUnlessDone(ctx, nil) }, ShouldNotPanic)
			})
		})
	})
	Convey("Given a cancelled context", t, func() {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		Convey("When an error occurs", func() {
			Convey("Then it is ignored and does not exit", func() {
				So(func() { fatal.OnErrorUnlessDone(ctx, errors.New("request cancelled")) }, ShouldNotPanic)
			})
		})
	})
}

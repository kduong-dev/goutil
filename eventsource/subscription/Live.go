package subscription

import (
	"context"
	"time"

	"github.com/kduong-dev/goutil/eventsource"
	"github.com/kduong-dev/goutil/fatal"
)

func Live(ctx context.Context, input Input) (cursor int64, err error) {
	const limit = 1000
	const perReadTimeout = 10 * time.Second
	var events []*eventsource.Event
	cursor = input.Cursor
	for {
		if err = ctx.Err(); err != nil {
			return
		}
		readCtx, cancel := context.WithTimeout(ctx, perReadTimeout)
		events, cursor, err = input.Log.Read(readCtx, cursor, limit)
		cancel()
		switch err {
		case nil:
			for _, event := range events {
				if err = input.Apply(ctx, event); err != nil {
					return
				}
			}
		case eventsource.Timeout:
			continue
		default:
			fatal.OnError(err)
		}
	}
}

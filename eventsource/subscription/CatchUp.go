package subscription

import (
	"context"

	"github.com/kduong-dev/goutil/eventsource"
	"github.com/kduong-dev/goutil/fatal"
)

func CatchUp(ctx context.Context, input Input) (cursor int64, err error) {
	const limit = 1000
	// Each read must return immediately rather than block, so CatchUp stops
	// as soon as it runs out of backlog.
	pollCtx, cancel := context.WithCancel(ctx)
	cancel()
	var events []*eventsource.Event
	cursor = input.Cursor
	for {
		events, cursor, err = input.Log.Read(pollCtx, cursor, limit)
		switch err {
		case nil:
			if len(events) == 0 {
				return
			}
			for _, event := range events {
				if err = input.Apply(ctx, event); err != nil {
					return
				}
			}
		case eventsource.Timeout:
			err = nil
			return
		default:
			fatal.OnError(err)
		}
	}
}

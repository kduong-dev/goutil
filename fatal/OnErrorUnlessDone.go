package fatal

import "context"

// OnErrorUnlessDone is OnError, except errors are ignored once ctx is done,
// since they are then typically caused by the cancellation itself.
//
// Deprecated: ctx is not reliably done when a client disconnects, and write errors after a response has
// started are I/O failures rather than bugs, so ignore or log them instead.
func OnErrorUnlessDone(ctx context.Context, err error, args ...any) {
	if err == nil || ctx.Err() != nil {
		return
	}
	OnError(err, args...)
}

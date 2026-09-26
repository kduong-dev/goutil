package fatal

import "context"

// OnErrorUnlessDone is OnError, except errors are ignored once ctx is done,
// since they are then typically caused by the cancellation itself.
func OnErrorUnlessDone(ctx context.Context, err error, args ...any) {
	if err == nil || ctx.Err() != nil {
		return
	}
	OnError(err, args...)
}

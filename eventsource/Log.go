package eventsource

import (
	"context"
	"errors"
	"io"
)

var (
	Timeout       = errors.New("timeout") // the log is empty
	UnknownCursor = errors.New("unknown cursor")
)

type Log interface {
	// Close releases resources held by the log implementation.
	// In-memory logs may implement this as a no-op, while network-backed logs
	// should close underlying clients/connections.
	io.Closer

	// Channel returns the log identifier used by the backend implementation.
	Channel() string

	// Append writes a new event payload and returns the created event.
	Append(data []byte) (event *Event, err error)

	// Read returns up to limit events with sequence > cursor. The returned
	// nextCursor is the highest sequence observed. Implementations may block
	// until events are available or ctx is done.
	Read(ctx context.Context, cursor int64, limit int) (events []*Event, nextCursor int64, err error)
}

type Event struct {
	LogID    string `json:"log_id"`
	Sequence int64  `json:"sequence"`
	Data     []byte `json:"data"`
}

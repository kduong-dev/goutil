package eventsource

import (
	"context"
	"sync"
)

type InMemoryLog struct {
	channel string

	mu     sync.Mutex
	events []*Event
	waitCh chan struct{}
}

func NewInMemoryLog(channel string) *InMemoryLog {
	return &InMemoryLog{
		channel: channel,
		waitCh:  make(chan struct{}),
	}
}

func (log *InMemoryLog) Close() error {
	return nil
}

func (log *InMemoryLog) Channel() string {
	return log.channel
}

func (log *InMemoryLog) Append(data []byte) (event *Event, err error) {
	log.mu.Lock()
	defer log.mu.Unlock()

	sequence := int64(len(log.events) + 1)
	event = &Event{
		LogID:    log.channel,
		Sequence: sequence,
		Data:     cloneBytes(data),
	}
	log.events = append(log.events, event)

	close(log.waitCh)
	log.waitCh = make(chan struct{})
	return event, nil
}

// Read returns up to limit events after cursor. If no events are
// immediately available, it blocks until events arrive or ctx is done. If
// ctx is already done when no events are immediately available, it returns
// with no error rather than blocking.
func (log *InMemoryLog) Read(ctx context.Context, cursor int64, limit int) (events []*Event, nextCursor int64, err error) {
	if limit <= 0 {
		return nil, cursor, nil
	}

	for {
		log.mu.Lock()
		available := log.availableEvents(cursor, limit)
		if len(available) > 0 {
			nextCursor = cursor
			events = make([]*Event, 0, len(available))
			for _, event := range available {
				copyEvent := &Event{
					LogID:    event.LogID,
					Sequence: event.Sequence,
					Data:     cloneBytes(event.Data),
				}
				events = append(events, copyEvent)
				if copyEvent.Sequence > nextCursor {
					nextCursor = copyEvent.Sequence
				}
			}
			log.mu.Unlock()
			return events, nextCursor, nil
		}

		if ctx.Err() != nil {
			log.mu.Unlock()
			return nil, cursor, nil
		}

		waitCh := log.waitCh
		log.mu.Unlock()

		select {
		case <-waitCh:
		case <-ctx.Done():
			return nil, cursor, Timeout
		}
	}
}

func (log *InMemoryLog) availableEvents(cursor int64, limit int) []*Event {
	start := cursor
	if start < 0 {
		start = 0
	}
	if start >= int64(len(log.events)) {
		return nil
	}

	end := start + int64(limit)
	if end > int64(len(log.events)) {
		end = int64(len(log.events))
	}
	return log.events[start:end]
}

func cloneBytes(data []byte) []byte {
	if data == nil {
		return nil
	}
	out := make([]byte, len(data))
	copy(out, data)
	return out
}

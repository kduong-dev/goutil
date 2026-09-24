package httpx

import (
	"errors"

	"github.com/ansel1/merry"
	"github.com/kduong-dev/goutil/fatal"
)

// MerrifiedSentinel pairs a sentinel error with the status code and user
// message it is sent to clients as.
type MerrifiedSentinel struct {
	Sentinel    error
	StatusCode  int
	UserMessage string
}

// MerrifiedSentinels are the sentinel errors a client can cause.
type MerrifiedSentinels []MerrifiedSentinel

func (merrifiedSentinels MerrifiedSentinels) lookup(err error) error {
	for _, merrified := range merrifiedSentinels {
		if errors.Is(err, merrified.Sentinel) {
			return merry.Wrap(err).WithHTTPCode(merrified.StatusCode).WithUserMessage(merrified.UserMessage)
		}
	}
	return nil
}

// Merrify attaches the status code and user message of the matching sentinel
// to err, and otherwise returns it unchanged.
func (merrifiedSentinels MerrifiedSentinels) Merrify(err error) error {
	if merrifiedError := merrifiedSentinels.lookup(err); merrifiedError != nil {
		return merrifiedError
	}
	return err
}

// MerrifyOrFatal is Merrify for calls whose only expected failures are the
// sentinels; any other error is fatal.
func (merrifiedSentinels MerrifiedSentinels) MerrifyOrFatal(err error) error {
	merrifiedError := merrifiedSentinels.lookup(err)
	if merrifiedError == nil {
		fatal.OnError(err)
	}
	return merrifiedError
}

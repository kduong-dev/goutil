package httpx

import (
	"errors"
	"net/http"

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
//
// Deprecated: a handler can't tell a bug from an I/O failure such as a timeout or a client
// disconnecting, and exiting on those takes every other request down with it. The package
// returning the error calls fatal.OnError for its own bugs; send everything else with
// SendErrorResponse, which responds 500 to errors that match no sentinel.
func (merrifiedSentinels MerrifiedSentinels) MerrifyOrFatal(err error) error {
	merrifiedError := merrifiedSentinels.lookup(err)
	if merrifiedError == nil {
		fatal.OnError(err)
	}
	return merrifiedError
}

// SendErrorResponse is SendErrorResponse for err merrified with the matching
// sentinel, so handlers can return the errors of the packages they call as is.
func (merrifiedSentinels MerrifiedSentinels) SendErrorResponse(responseWriter http.ResponseWriter, err error) {
	SendErrorResponse(responseWriter, merrifiedSentinels.Merrify(err))
}

package httpx

import (
	"net/http"

	"github.com/ansel1/merry"
)

// UpstreamResponseError is ResponseError for responses from third-party
// APIs. The upstream status code is replaced with 502 Bad Gateway, since
// codes such as 401 describe our credentials with the upstream, not the
// client's with us.
func UpstreamResponseError(response *http.Response) error {
	err := ResponseError(response)
	if err == nil {
		return nil
	}
	return merry.Wrap(err).WithHTTPCode(http.StatusBadGateway)
}

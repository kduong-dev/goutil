package httpx

import "io"

// DrainAndClose discards the rest of body and closes it, so that the
// http.Client can reuse the underlying connection.
func DrainAndClose(body io.ReadCloser) error {
	_, err := io.Copy(io.Discard, body)
	if err != nil {
		body.Close()
		return err
	}
	return body.Close()
}

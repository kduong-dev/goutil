package httpx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ansel1/merry"
)

// ResponseError returns nil unless response's status code is >= 400.
// Otherwise it returns an error carrying response's HTTP status code. The
// user message is taken from the response body's ResponseMessage JSON (as
// written by SendErrorResponse) when present, falling back to the
// response's status text; the raw body is never used as the user message,
// since it may come from a third-party API and isn't safe to forward to
// our own clients. The raw body is kept in the error's own message for
// logging. ResponseError panics if the response body cannot be read.
func ResponseError(response *http.Response) error {
	if response.StatusCode < http.StatusBadRequest {
		return nil
	}
	userMessage := http.StatusText(response.StatusCode)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var responseMessage ResponseMessage
	err = json.Unmarshal(body, &responseMessage)
	containsResponseMessage := err == nil && responseMessage.Message != ""
	if containsResponseMessage {
		userMessage = responseMessage.Message
	}
	message := fmt.Sprintf("request failed with status %d: %s", response.StatusCode, strings.TrimSpace(string(body)))
	return merry.New(message).WithHTTPCode(response.StatusCode).WithUserMessage(userMessage)
}

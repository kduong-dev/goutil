package httpx

import (
	"encoding/json"
	"net/http"
)

// SendJSONResponse marshals before writing the status so a marshal bug never sends a half-written body, and
// ignores write errors since they only mean the client has gone.
func SendJSONResponse(responseWriter http.ResponseWriter, statusCode int, body any) {
	responseWriter.Header().Set("Content-Type", "application/json")
	if body == nil {
		responseWriter.WriteHeader(statusCode)
		return
	}
	encoded, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	responseWriter.WriteHeader(statusCode)
	_, _ = responseWriter.Write(append(encoded, '\n'))
}

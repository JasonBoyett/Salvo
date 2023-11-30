package runner

import (
	"net/http"
	"time"
)

type Result struct {
	// Start is the time the request was started
	Start time.Time
	// End is the time the request was completed
	End time.Time
	// Success is true if the request was successful
	Success bool
	// The protocol version for incoming server requests.
	Proto string
	// Code is the HTTP status code of the response
	Code int
	// The response body
	Body   string
	Header http.Header
}

func (r Result) Duration() time.Duration {
	return r.End.Sub(r.Start)
}

func CountFailures(results []Result) int {
	fails := 0
	for _, result := range results {
		if !result.Success {
			fails++
		}
	}
	return fails
}

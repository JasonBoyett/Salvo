package runner

import (
	"time"
)

type Result struct {
	Start        time.Time `json:"start"`
	End          time.Time `json:"end"`
	Success      bool      `json:"success"`
	StatusCode   int       `json:"statusCode"`
	ResponseBody string    `json:"responseBody"`
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

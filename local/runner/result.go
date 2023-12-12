package runner

import (
	"time"
)

type Result struct {
	Start time.Time
	End time.Time
	Success bool
	Code int
	Body string
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

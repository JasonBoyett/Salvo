package runner

import (
	"time"
)

type Opts struct {
	Path string        `json:"path"`
	Time time.Duration `json:"time"`
	// The time between requests
	Users        int     `json:"users"`
	Timeout      int     `json:"timeout"`
	SuccessCodes []int   `json:"successCodes"`
	Rate         float32 `json:"rate"`
	// If Rate is 0, the requests will be made as fast as possible
}

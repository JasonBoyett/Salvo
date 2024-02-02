package runner

import (
	"time"
)

type Opts struct {
	Path         string
	Time         time.Duration
	Users        int
	Timeout      int
	SuccessCodes []int
	Rate         *float64 // If Rate is nil, the requests will be made as fast as possible
	ResultBody   string
}

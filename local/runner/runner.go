package runner

import (
	"net/http"
	"sync"
	"time"
)

// Opts is a struct that contains the options for the runner.
type Opts struct {
  // Path is the path to make the request to
	Path    string
  // Time is the time in seconds to run the test for
	Time    int
  // Users is the number of users to simulate
	Users   int
  // Timeout is the timeout in seconds for each request
	Timeout int
  // SuccessCodes is a slice of success codes
  SuccessCodes []int
  // Rate is the rate in requests per second. If nil,
  // the requests will be made as fast as possible
	Rate    *float64 //Rate is a pointer so that it can be nil
}

// Run executes the load test with the given options.
// It returns a slice of results and the number of failed requests.
//
// The load test is configured using the provided options (Opts).
//
// Parameters:
//   - opts: Opts
//     The options for the load test.
//     - Path: string
//       The path to make the request to.
//     - Time: int
//       The duration in seconds to run the test for.
//     - Users: int
//       The number of users to simulate.
//     - Timeout: int
//       The timeout in seconds for each request.
//     - Rate: *float64
//       The rate in requests per second. 
//       If nil, the requests will be made as fast as possible.
func Run(opts Opts) ([]Result, int){

  results := make([]Result, 0)
  fails := 0


	failsCh := make(chan int)
	resultsCh := make(chan Result)

	wg := new(sync.WaitGroup)

  for i := 0; i < opts.Users; i++ {

  go simUser(opts, wg, failsCh, resultsCh)

  }

  
  wg.Wait()
  for value := range failsCh {
    fails += value
  }
  close(failsCh)
  for value := range resultsCh {
    results = append(results, value)
  }
  close(resultsCh)

  return results, fails
}

func simUser(opts Opts, wg *sync.WaitGroup, failsCh chan int, resultsCh chan Result) {

	milisecondsPerSecond := 1_000

  timerStart := time.Now()

  for {

    start := time.Now()
    wg.Add(1)
    requestGood := true

    status, err := makeRequest(opts.Path, opts.Timeout)
    if err != nil || status != http.StatusOK || !contains(opts.SuccessCodes, status) {
      failsCh <- 1
      requestGood = false
    }

    resultsCh <- Result{
      start,
      time.Now(),
      requestGood,
      status,
    }

    if opts.Rate != nil {

      rate := *opts.Rate
      time.Sleep(time.Duration(float64(milisecondsPerSecond) / rate) * time.Millisecond)

    }

    if time.Since(timerStart).Seconds() > float64(opts.Time) {
      wg.Done()
      break
    }
  }
}

// makeRequest makes a GET request to the given path with a specified timeout.
//
// Parameters:
//   - path: string
//     The path to make the request to.
//   - timeout: int
//     The timeout in seconds.
//
// Returns:
//   - int
//     The status code of the HTTP response.
//   - error
//     An error if one occurred during the request.
func makeRequest(path string, timeout int) (int, error) {
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}

	response, err := client.Get(path)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	defer response.Body.Close()
	return response.StatusCode, nil
}

// contains checks if a slice of integers contains a given integer.
func contains(slice []int, value int) bool {
  for _, item := range slice {
    if item == value {
      return true
    }
  }
  return false
}

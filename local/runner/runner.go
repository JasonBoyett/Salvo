package runner

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Opts is a struct that contains the options for the runner.
type Opts struct {
	// Path is the path to make the request to
	Path string
	// Time is the time in seconds to run the test for
	Time int
	// Users is the number of users to simulate
	Users int
	// Timeout is the timeout in seconds for each request
	Timeout int
	// SuccessCodes is a slice of success codes
	SuccessCodes []int
	// Rate is the rate in requests per second. If nil,
	// the requests will be made as fast as possible
	Rate *float64 //Rate is a pointer so that it can be nil
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
func Run(opts Opts) ([]Result, int) {

	results := make([]Result, 0)
	fails := 0

	failsCh := make(chan int)
	resultsCh := make(chan Result)

	var wg sync.WaitGroup

	for i := 0; i < opts.Users; i++ {

		wg.Add(1)
		fmt.Println("starting user")
		go simUser(opts, &wg, failsCh, resultsCh)

	}

	// go func() {
	// 	wg.Wait()
	// 	close(failsCh)
	// 	close(resultsCh)
	// }()

	for value := range failsCh {
		fails += value
	}
	for value := range resultsCh {
		results = append(results, value)
	}

	return results, fails
}

func simUser(opts Opts, wg *sync.WaitGroup, failsCh chan<- int, resultsCh chan<- Result) {

	defer wg.Done()
	fmt.Println("simulating user")
	startTime := time.Now()

  for {

		elapsedTime := time.Since(startTime).Seconds()
		if elapsedTime >= float64(opts.Time) {
			break
		}
     
		fmt.Println("starting loop")
		start := time.Now()
		result, err := makeRequest(opts.Path, opts.Timeout)
		if err != nil {

			fmt.Println("writing error")
			failsCh <- 1

			resultsCh <- Result{
				Start:   start,
				End:     time.Now(),
				Success: false,
				Code:    result,
			}
      fmt.Println("error written")

		} else {

			fmt.Println("writing success")
			resultsCh <- Result{
				Start:   start,
				End:     time.Now(),
				Success: true,
				Code:    result,
			}
      fmt.Println("success written")

		}

		if opts.Rate != nil {
			fmt.Println("sleeping")
			time.Sleep(time.Duration(1 / *opts.Rate) * time.Second)
		}
		fmt.Println("loop done")
	}
	fmt.Println("user done")
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

	fmt.Println("makiing call")
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}

	response, err := client.Get(path)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	defer response.Body.Close()
	fmt.Println("request closed")
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

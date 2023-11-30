package runner

import (
	"io"
	"net/http"
	"sync"
	"time"

	"golang.org/x/exp/slices"
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

// Run executes the load test with the given options.
// It returns a slice of results and the number of failed requests.
//
// The load test is configured using the provided options (Opts).
//
// Parameters:
//   - opts: Opts
//     The options for the load test.
//   - Path: string
//     The path to make the request to.
//   - Time: int
//     The duration in seconds to run the test for.
//   - Users: int
//     The number of users to simulate.
//   - Timeout: int
//     The timeout in seconds for each request.
//   - Rate: *float64
//     The rate in requests per second.
//     If nil, the requests will be made as fast as possible.
func Run(opts Opts) ([]Result, int) {
	results := make([]Result, 0)
	fails := 0

	failsCh := make(chan int)
	resultsCh := make(chan Result)

	var wg sync.WaitGroup
	var failsGroup sync.WaitGroup
	var resultsGroup sync.WaitGroup

	defer failsGroup.Wait()
	defer resultsGroup.Wait()

	for i := 0; i < opts.Users; i++ {

		wg.Add(1)
		go simUser(opts, &wg, failsCh, resultsCh)

	}

	go func() {
		wg.Wait()
		close(failsCh)
		close(resultsCh)
	}()

	failsGroup.Add(1)
	go func() {
		for value := range failsCh {
			fails += value
		}
		failsGroup.Done()
	}()

	resultsGroup.Add(1)
	go func() {
		for value := range resultsCh {
			results = append(results, value)
		}
		resultsGroup.Done()
	}()

	resultsGroup.Wait()
	failsGroup.Wait()
	return results, fails
}

func simUser(opts Opts, wg *sync.WaitGroup, failsCh chan<- int, resultsCh chan<- Result) {
	defer wg.Done()
	startTime := time.Now()

	for {

		elapsedTime := time.Since(startTime).Seconds()
		if elapsedTime >= float64(opts.Time) {
			break
		}

		start := time.Now()
		response, err := makeRequest(opts.Path, opts.Timeout)
		responseCode := response.code
		responseBody := response.body
		if err != nil || !slices.Contains(opts.SuccessCodes, responseCode) || responseCode != http.StatusOK {

			failsCh <- 1

			resultsCh <- Result{
				Start:   start,
				End:     time.Now(),
				Success: false,
				Code:    responseCode,
				Body:    responseBody,
			}

		} else {
			resultsCh <- Result{
				Start:   start,
				End:     time.Now(),
				Success: true,
				Code:    responseCode,
				Body:    responseBody,
			}
		}

		if opts.Rate != nil {
			time.Sleep(time.Duration(1 / *opts.Rate) * time.Second)
		}
	}
}

// Contains the completed response
//
// Body is a string of the enitre response body
// Rather than a io.ReadCloser
type finalResponse struct {
	code   int
	body   string
	header http.Header
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
//   - finishedResponse
//     A struct containing the response
//   - error
//     An error if one occurred during the request.
func makeRequest(path string, timeout int) (finalResponse, error) {
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}

	result := finalResponse{
		code: http.StatusInternalServerError,
		body: "",
	}
	response, err := client.Get(path)
	if err != nil {
		return result, err
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return result, err
	}

	result = finalResponse{
		code:   response.StatusCode,
		body:   string(responseBody),
		header: response.Header,
	}

	return result, nil
}

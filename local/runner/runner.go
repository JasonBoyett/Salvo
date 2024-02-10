package runner

import (
	"io"
	"net/http"
	"sync"
	"time"

	"golang.org/x/exp/slices"
)

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
//   - Rate: float64
//     The rate in requests per second.
//     If Rate is 0, the requests will be made as fast as possible.
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

func simUser(
	opts Opts, wg *sync.WaitGroup,
	failsCh chan<- int,
	resultsCh chan<- Result,
) {
	defer wg.Done()
	startTime := time.Now()

	for {
		elapsedTime := time.Since(startTime).Seconds()
		if elapsedTime >= float64(opts.Time) {
			break
		}

		start := time.Now()
		response, err := MakeRequest(opts.Path, opts.Timeout)
		responseCode := response.Code
		responseBody := response.Body
		if err != nil ||
			responseCode != http.StatusOK {

			failsCh <- 1

			resultsCh <- Result{
				Start:        start,
				End:          time.Now(),
				Success:      false,
				StatusCode:   responseCode,
				ResponseBody: responseBody,
			}

		} else if responseCode != http.StatusOK {

			failsCh <- 1

			resultsCh <- Result{
				Start:        start,
				End:          time.Now(),
				Success:      false,
				StatusCode:   responseCode,
				ResponseBody: responseBody,
			}

		} else if !slices.Contains(opts.SuccessCodes, responseCode) &&
			opts.SuccessCodes != nil {

			failsCh <- 1

			resultsCh <- Result{
				Start:        start,
				End:          time.Now(),
				Success:      false,
				StatusCode:   responseCode,
				ResponseBody: responseBody,
			}

		} else {
			resultsCh <- Result{
				Start:        start,
				End:          time.Now(),
				Success:      true,
				StatusCode:   responseCode,
				ResponseBody: responseBody,
			}
		}

		if opts.Rate != 0 {
			time.Sleep(time.Duration(1 / opts.Rate) * time.Second)
		}
	}
}

// Contains the completed response
//
// Body is a string of the entire response body
// Rather than a io.ReadCloser
type FinalResponse struct {
	Code int
	Body string
}

// MakeRequest makes a GET request to the given path with a specified timeout.
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
func MakeRequest(path string, timeout int) (FinalResponse, error) {
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}

	result := FinalResponse{
		Code: http.StatusInternalServerError,
		Body: "",
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

	result = FinalResponse{
		Code: response.StatusCode,
		Body: string(responseBody),
	}

	return result, nil
}

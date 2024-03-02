package test

import (
	"github.com/JasonBoyett/salvo/local/runner"
	"testing"
	"time"

	server "github.com/JasonBoyett/salvo/local/runner/server"
)

func TestCall(t *testing.T) {
	response, err := runner.MakeRequest("https://random-word-api.herokuapp.com/languages", 10)
	if err != nil {
		t.Errorf("Error making request: %s", err)
	}

	code := response.Code

	if code != 200 {
		t.Fatalf("Expected 200, got %d", code)
	}
}

func TestCallServer(t *testing.T) {
	route := "/"

	serverOpts := server.TestOpts{
		Message: "test1",
		Fail:    false,
		Delay:   1 * time.Second,
		Port:    8085,
		Route:   route,
	}

	err := server.TestServer(serverOpts)
	if err != nil {
		t.Errorf("Error starting server: %s", err)
	}
	response, err := runner.MakeRequest("http://localhost:8085"+route, 10)
	code := response.Code
	time.Sleep(10 * time.Second)
	if err != nil {
		t.Errorf("Error making request: %s", err)
	}
	if code != 200 {
		t.Fatalf("Expected 200, got %d", code)
	}
}

func TestRunner(t *testing.T) {

	route := "/test2"

	serverOpts := server.TestOpts{
		Message: "hello",
		Fail:    false,
		Delay:   1 * time.Second,
		Port:    8086,
		Route:   route,
	}

	testOpts := runner.Opts{
		Path:         "http://localhost:8086" + route,
		Time:         10,
		Users:        10,
		Timeout:      10,
		SuccessCodes: []int{200},
		Rate:         1.1,
	}

	err := server.TestServer(serverOpts)
	if err != nil {
		t.Errorf("Error starting server: %s", err)
	}

	results, fails := runner.Run(testOpts)

	if fails != 0 {
		t.Fatalf("Expected 0 fails, got %d", fails)
	}

	if len(results) != 100 {
		t.Fatalf("Expected 100 results, got %d", len(results))
	}
}

func TestCallingFailingServer(t *testing.T) {

	route := "/test3"

	serverOpts := server.TestOpts{
		Message: "you're a failure",
		Fail:    true,
		Delay:   1 * time.Second,
		Port:    8087,
		Route:   route,
	}

	testOpts := runner.Opts{
		Path:         "http://localhost:8087" + route,
		Time:         10,
		Users:        10,
		Timeout:      10,
		SuccessCodes: []int{200},
		Rate:         1.1,
	}

	err := server.TestServer(serverOpts)
	if err != nil {
		t.Errorf("Error starting server: %s", err)
	}

	results, _ := runner.Run(testOpts)

	if runner.CountFailures(results) != 100 {
		t.Fatalf("Expected 100 fails, got %d", runner.CountFailures(results))
	}

	if len(results) != 100 {
		t.Fatalf("Expected 100 results, got %d", len(results))
	}
}

func TestResponseBody(t *testing.T) {

	route := "/test4"

	serverOpts := server.TestOpts{
		Message: "hello",
		Fail:    false,
		Delay:   1 * time.Second,
		Port:    8088,
		Route:   route,
	}

	testOpts := runner.Opts{
		Path:         "http://localhost:8088" + route,
		Time:         2,
		Users:        1,
		Timeout:      10,
		SuccessCodes: []int{200},
		Rate:         1.1,
	}

	err := server.TestServer(serverOpts)
	if err != nil {
		t.Errorf("Error starting server: %s", err)
	}

	results, _ := runner.Run(testOpts)

	for _, result := range results {
		if result.ResponseBody != "hello" {
			t.Fatalf("Expected \"hello\" but got \"%s\"", result.ResponseBody)
		}
	}
}

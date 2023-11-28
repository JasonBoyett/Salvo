package runner

import (
	"testing"
	"time"

	server "github.com/JasonBoyett/spyglass/local/runner/server"
)

func TestCall(t *testing.T) {
	response, err := makeRequest("https://random-word-api.herokuapp.com/languages", 10)
	if err != nil {
		t.Errorf("Error making request: %s", err)
	}

	// Testing the responde code
	code := response.code

	if code != 200 {
		t.Errorf("Expected 200, got %d", code)
	}

	// Testing the response body
	body := response.body
	if body != `["es","zh","it","de"]` {
		t.Errorf(`Expected ["es","zh","it","de"], got %s`, body)
	}
}

func TestCallServer(t *testing.T) {
	root := "/"

	serverOpts := server.TestOpts{
		Message: "test1",
		Fail:    false,
		Delay:   1 * time.Second,
		Port:    8085,
		Rout:    root,
	}

	err := server.TestServer(serverOpts)
	if err != nil {
		t.Errorf("Error starting server: %s", err)
	}
	response, err := makeRequest("http://localhost:8085"+root, 10)
	code := response.code
	time.Sleep(10 * time.Second)
	if err != nil {
		t.Errorf("Error making request: %s", err)
	}
	if code != 200 {
		t.Errorf("Expected 200, got %d", code)
	}
}

func TestRunner(t *testing.T) {
	rout := "/test2"
	rate := 1.1

	serverOpts := server.TestOpts{
		Message: "hello",
		Fail:    false,
		Delay:   1 * time.Second,
		Port:    8086,
		Rout:    rout,
	}

	testOpts := Opts{
		Path:         "http://localhost:8086" + rout,
		Time:         10,
		Users:        10,
		Timeout:      10,
		SuccessCodes: []int{200},
		Rate:         &rate,
	}

	err := server.TestServer(serverOpts)
	if err != nil {
		t.Errorf("Error starting server: %s", err)
	}

	results, fails := Run(testOpts)

	if fails != 0 {
		t.Errorf("Expected 0 fails, got %d", fails)
	}

	if len(results) != 100 {
		t.Errorf("Expected 100 results, got %d", len(results))
	}
}

func TestCallingFailingServer(t *testing.T) {
	rout := "/test3"
	rate := 1.1

	serverOpts := server.TestOpts{
		Message: "you're a failure",
		Fail:    true,
		Delay:   1 * time.Second,
		Port:    8087,
		Rout:    rout,
	}

	testOpts := Opts{
		Path:         "http://localhost:8087" + rout,
		Time:         10,
		Users:        10,
		Timeout:      10,
		SuccessCodes: []int{200},
		Rate:         &rate,
	}

	err := server.TestServer(serverOpts)
	if err != nil {
		t.Errorf("Error starting server: %s", err)
	}

	results, _ := Run(testOpts)

	if CountFailures(results) != 100 {
		t.Errorf("Expected 100 fails, got %d", CountFailures(results))
	}

	if len(results) != 100 {
		t.Errorf("Expected 100 results, got %d", len(results))
	}
}

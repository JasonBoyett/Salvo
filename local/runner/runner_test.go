package runner

import (
	"fmt"
	"testing"
	"time"

	server "github.com/JasonBoyett/spyglass/local/runner/server"
)

func TestCall(t *testing.T) {
	code, err := makeRequest("https://random-word-api.herokuapp.com/languages", 10)
	if err != nil {
		t.Errorf("Error making request: %s", err)
	}
	if code != 200 {
		t.Errorf("Expected 200, got %d", code)
	}
}

// func TestCallServer(t *testing.T) {
//   serverOpts := server.TestOpts{
//     Message: "hello",
//     Fail:    false,
//     Delay:   1 * time.Second,
//     Port:    8085,
//   }
//
//   err := server.TestServer(serverOpts)
//
//   if err != nil {
//     t.Errorf("Error starting server: %s", err)
//   }
//   code, err := makeRequest("http://localhost:8085/hello", 10)
//   time.Sleep(10 * time.Second)
//   if err != nil {
//     t.Errorf("Error making request: %s", err)
//   }
//   if code != 200 {
//     t.Errorf("Expected 200, got %d", code)
//   } else {
//     fmt.Println(code)
//   }
// }

func TestFailingServer(t *testing.T) {
	fmt.Println("test running")
	rate := 1.1

	serverOpts := server.TestOpts{
		Message: "hello",
		Fail:    false,
		Delay:   1 * time.Second,
		Port:    8086,
	}

	testOpts := Opts{
		Path:         "http://localhost:8086/hello",
		Time:         10,
		Users:        10,
		Timeout:      10,
		SuccessCodes: []int{200},
		Rate:         &rate,
	}

	err := server.TestServer(serverOpts)

	fmt.Println("server isn't blocking")

	if err != nil {
		t.Errorf("Error starting server: %s", err)
	}
	fmt.Println("server starting")

	results, fails := Run(testOpts)

	if fails != 0 {
		t.Errorf("Expected 0 fails, got %d", fails)
	}

	if len(results) != 100 {
    fmt.Println(results)
		t.Errorf("Expected 10 results, got %d", len(results))
	}
}

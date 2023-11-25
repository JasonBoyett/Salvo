package runner

import (
  server "github.com/JasonBoyett/spyglass/local/runner/server"
  "testing"
  "time"
)

func TestRunner(t *testing.T) {
  rate := 1.1
  serverOpts := server.TestOpts{
    Message: "hello",
    Fail: false,
    Delay: 1 * time.Second,
    Port:  8080,
  }

  testOpts := Opts{
    Path: "http://localhost:8080",
    Users: 10,
    Timeout: 1,
    Rate: &rate,
  }

  kill, err := server.TestServer(serverOpts)
  if err != nil {
    t.Fatal(err)
  }
  
  results, fails := Run(testOpts)
  
  if fails != 0 {
    t.Errorf("Expected 0 fails, got %d", fails)
  }

  if len(results) != 10 {
    t.Errorf("Expected 10 results, got %d", len(results))
  }

  err = kill()
  if err != nil {
    t.Fatal(err)
  }
}




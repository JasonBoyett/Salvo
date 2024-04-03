package test

import (
	"os/exec"
	"strings"
	"testing"
)

// This test requires that salvo has been built and is in the stystem path
// This test only tests weather the cli returns properly formatted json
// It does not test the accuracy of the json
// other tests in the suite cover the accuracy of salvo's json output
func TestCliOutputJson(t *testing.T) {
	jsonString := `
  {
    "path":"http://localhost:5081/TestCli",
    "time":1,
    "users":1,
    "timeout":10,
    "successCodes": [200],
    "rate":1.0
  }
  `
	args := []string{"-json", "-opts=" + jsonString}

	bytes, err := captureStdOut(args)
	result := string(bytes)
	if err != nil {
		t.Errorf("Error capturing stdout: %s", err)
	}

	if !strings.Contains(result, `"results":`) {
		t.Fatalf("Expected Results to be in output\n%s", result)
	}
	if !strings.Contains(result, `"start":`) {
		t.Fatalf("Expected Start to be in output\n%s", result)
	}
	if !strings.Contains(result, `"end":`) {
		t.Fatalf("Expected End to be in output\n%s", result)
	}
	if !strings.Contains(result, `"success":`) {
		t.Fatalf("Expected Success to be in output\n%s", result)
	}
	if !strings.Contains(result, `"statusCode":`) {
		t.Fatalf("Expected StatusCode to be in output\n%s", result)
	}
	if !strings.Contains(result, `"responseBody":`) {
		t.Fatalf("Expected ResponseBody to be in output\n%s", result)
	}
	if !strings.Contains(result, `"fails":`) {
		t.Fatalf("Expected Fails to be in output\n%s", result)
	}
}

func captureStdOut(args []string) ([]byte, error) {
	cmd := exec.Command("salvo", args...)
	stdout, err := cmd.Output()
	if err != nil {
		return []byte{}, err
	}
	return stdout, nil
}

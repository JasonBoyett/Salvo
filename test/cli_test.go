package test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	salvo "github.com/JasonBoyett/salvo/public"
)

func TestCliOutputFormat(t *testing.T) {
	originalArgs := os.Args
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
	args := []string{"salvo", "-json", "-opts=" + jsonString}

	result := captureStdOut(func() { salvo.Main() }, args)

	if !strings.Contains(result, `"results":`) {
		t.Errorf("Expected Results to be in output\n%s", result)
	}
	if !strings.Contains(result, `"start":`) {
		t.Errorf("Expected Start to be in output\n%s", result)
	}
	if !strings.Contains(result, `"end":`) {
		t.Errorf("Expected End to be in output\n%s", result)
	}
	if !strings.Contains(result, `"success":`) {
		t.Errorf("Expected Success to be in output\n%s", result)
	}
	if !strings.Contains(result, `"statusCode":`) {
		t.Errorf("Expected StatusCode to be in output\n%s", result)
	}
	if !strings.Contains(result, `"responseBody":`) {
		t.Errorf("Expected ResponseBody to be in output\n%s", result)
	}
	if !strings.Contains(result, `"fails":`) {
		t.Errorf("Expected Fails to be in output\n%s", result)
	}

	t.Cleanup(func() {
		os.Args = originalArgs
		fmt.Println("Resetting os.Args")
	})
}

func captureStdOut(f func(), args []string) string {
	originalArgs := os.Args
	os.Args = args
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	os.Stdout = originalStdout
	os.Args = originalArgs
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

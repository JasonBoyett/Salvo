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
    "Path":"http://localhost:5081/TestCli"
    ,"Time":1,
    "Users":1,
    "Timeout":10,
    "SuccessCodes": [200],
    "Rate":1.0
  }
  `
  args := []string{"salvo", "-json", "-opts=" + jsonString}

	result := captureStdOut(func() { salvo.Main() }, args)

	if !strings.Contains(result, `"Results":`) {
		t.Errorf("Expected Results to be in output\n%s", result)
	}
	if !strings.Contains(result, `"Start":`) {
		t.Errorf("Expected Start to be in output\n%s", result)
	}
	if !strings.Contains(result, `"End":`) {
		t.Errorf("Expected End to be in output\n%s", result)
	}
	if !strings.Contains(result, `"Success":`) {
		t.Errorf("Expected Success to be in output\n%s", result)
	}
	if !strings.Contains(result, `"StatusCode":`) {
		t.Errorf("Expected StatusCode to be in output\n%s", result)
	}
	if !strings.Contains(result, `"ResponseBody":`) {
		t.Errorf("Expected ResponseBody to be in output\n%s", result)
	}
	if !strings.Contains(result, `"Fails":`) {
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

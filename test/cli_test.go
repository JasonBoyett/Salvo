package test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	salvo "github.com/JasonBoyett/salvo/public"
)

func TestCliOutputFormat(t *testing.T) {
  originalArgs := os.Args
  defer func() { os.Args = originalArgs }()
  
  jsonString := `
  {
    "Path":"http://localhost:7081/TestReceiverWithOutsideJson"
    ,"Time":1,
    "Users":1,
    "Timeout":10,
    "SuccessCodes": [200],
    "Rate":1.0
  }
  `
  os.Args = []string{"salvo", "-json", "-opts=" + jsonString}

  result := captureStdOut(func(){ salvo.Main()})
  
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
}

func captureStdOut(f func()) string {
  originalStdout := os.Stdout
  r, w, _ := os.Pipe()
  os.Stdout = w
  f()
  w.Close()
  os.Stdout = originalStdout
  var buf bytes.Buffer
  io.Copy(&buf, r)
  return buf.String()
}

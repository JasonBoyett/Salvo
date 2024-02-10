package test

import (
	"bytes"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/JasonBoyett/salvo/local/runner"
	server "github.com/JasonBoyett/salvo/local/runner/server"
	"github.com/JasonBoyett/salvo/public/protobuff"
	"google.golang.org/protobuf/proto"
)

//this global is here so we can compare the result one of the tests
// against another test case. This is not ideal, but since protobuffs
// are not human readable, we need to compare the bytes directly.
var globalMessageBytes []byte

func TestProtobuffReceiver(t *testing.T) {
	route := "/protobuff"
	port := "6080"
	protInt, _ := strconv.Atoi(port)

	testOpts := runner.Opts{
		Path:         "http://localhost:" + port + route,
		Time:         1,
		Users:        1,
		Timeout:      10,
		SuccessCodes: []int{200},
		Rate:         1.0,
	}

	serverOpts := server.TestOpts{
		Message: "json test1",
		Fail:    false,
		Delay:   1 * time.Second,
		Port:    protInt,
		Route:   route,
	}

	err := server.TestServer(serverOpts)
	if err != nil {
		t.Errorf("Error starting server: %s", err)
	}
	pbuff := protobuff.OptsToProtobuff(&testOpts)

	res := protobuff.Receive(pbuff)

	if res.Results == nil {
		t.Errorf("Error receiving data")
	}
}

func TestProtobuffSenderStdout(t *testing.T) {
	var out bytes.Buffer

	resultOne := runner.Result{
		Start:        time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		End:          time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		StatusCode:   200,
		Success:      true,
		ResponseBody: "test",
	}
	resultTwo := runner.Result{
		Start:        time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		End:          time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		StatusCode:   42,
		Success:      true,
		ResponseBody: "foo",
	}
	resultThree := runner.Result{
		Start:        time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		End:          time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		StatusCode:   23,
		Success:      true,
		ResponseBody: "bar",
	}
	message := protobuff.Message{
		Results: []runner.Result{resultOne, resultTwo, resultThree},
		Fails:    0,
	}

  err := protobuff.SendStdout(&out, message)
  if err != nil {
    t.Errorf("Error sending to stdout: %s", err)
  }
  formattedMessage := protobuff.FormatMessage(message)

  expected, err := proto.Marshal(formattedMessage)
  if err != nil {
    t.Errorf("Error marshaling data: %s", err)
  }

  if !bytes.Equal(out.Bytes(), expected) {
    t.Errorf("Expected %v, got %v", expected, out.Bytes())
  }
  globalMessageBytes = out.Bytes()
}

func TestProtobuffSenderToFile(t *testing.T) {
  // section 1 write protobuff to file
	resultOne := runner.Result{
		Start:        time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		End:          time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		StatusCode:   200,
		Success:      true,
		ResponseBody: "test",
	}
	resultTwo := runner.Result{
		Start:        time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		End:          time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		StatusCode:   42,
		Success:      true,
		ResponseBody: "foo",
	}
	resultThree := runner.Result{
		Start:        time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		End:          time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		StatusCode:   23,
		Success:      true,
		ResponseBody: "bar",
	}
	message := protobuff.Message{
		Results: []runner.Result{resultOne, resultTwo, resultThree},
		Fails:    0,
	}

  // fileName := "testfile.pb" 
  localPath := "./"
  info, err := os.Stat(localPath)
  if err != nil {
    t.Errorf("Error getting file info: %s", err)
  }
  path, err := filepath.Abs(info.Name())
  if err != nil {
    t.Errorf("Error getting absolute path: %s", err)
  }
  err = protobuff.SendFile(path, message)
  if err != nil {
    t.Errorf("Error sending to file: %s", err)
  }

  // section 2 read protobuff from file, test values, and delete file
	file, err := os.Open(path)
	if err != nil {
		t.Errorf("Error opening file: %s", err)
	}
	defer file.Close()

	protobuffPath, err := filepath.Abs(file.Name())
	if err != nil {
		t.Errorf("Error getting absolute path: %s", err)
	}
  protobuffPath = protobuffPath + string(filepath.Separator) + "salvo.pb"
	pbuff, err := os.ReadFile(protobuffPath)
	if err != nil {
		t.Errorf("Error reading file: %s", err)
	}

  if !bytes.Equal(pbuff, globalMessageBytes) {
    t.Errorf("Expected %v, got %v", globalMessageBytes, pbuff)
  }

	err = os.Remove(protobuffPath)
	if err != nil {
	  t.Errorf("Error removing file: %s", err)
	}
}

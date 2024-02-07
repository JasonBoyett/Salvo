package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/JasonBoyett/salvo/local/runner"
	server "github.com/JasonBoyett/salvo/local/runner/server"
	salvoJson "github.com/JasonBoyett/salvo/public/json"
)

func TestJsonReceiver(t *testing.T) {
	route := "/json"
	port := "7080"
	protInt, _ := strconv.Atoi(port)
	rate := 1.0

	testOpts := runner.Opts{
		Path:         "http://localhost:" + port + route,
		Time:         1,
		Users:        1,
		Timeout:      10,
		SuccessCodes: []int{200},
		Rate:         &rate,
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
	json, err := json.Marshal(testOpts)
	if err != nil {
		t.Errorf("Error marshaling data: %s", err)
	}
	message, err := salvoJson.Receive(json)
	if err != nil {
		t.Errorf("Error receiving data: %s", err)
	}
	if message.Fails != 0 {
		t.Errorf("Expected 0 fails, got %d", message.Fails)
	}

	result := message.Results[0]
	if result.ResponseBody != "json test1" {
		t.Errorf("Expected json test1, got %s", result.ResponseBody)
	}
}

func TestReciverWithOutsideJson(t *testing.T) {

	serverOpts := server.TestOpts{
		Message: "json test2",
		Fail:    false,
		Delay:   1 * time.Second,
		Port:    7081,
		Route:   "/TestReceiverWithOutsideJson",
	}

	path := "test_receiver.json"
	json, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("Error reading file: %s", err)
	}

	err = server.TestServer(serverOpts)
	if err != nil {
		t.Errorf("Error starting server: %s", err)
	}
	if err != nil {
		t.Errorf("Error marshaling data: %s", err)
	}
	message, err := salvoJson.Receive(json)
	if err != nil {
		t.Errorf("Error receiving data: %s", err)
	}
	if message.Fails != 0 {
		t.Errorf("Expected 0 fails, got %d", message.Fails)
	}

	result := message.Results[0]
	if result.ResponseBody != "json test2" {
		t.Errorf("Expected json test2, got %s", result.ResponseBody)
	}
}

func TestJsonSenderStdout(t *testing.T) {
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
	message := salvoJson.Message{
		Results: []runner.Result{resultOne, resultTwo, resultThree},
		Fails:   0,
	}
	err := salvoJson.SendStdout(&out, message)
	if err != nil {
		t.Errorf("Error sending to stdout: %s", err)
	}

	if strings.Contains(out.String(), `"ResponseBody":"test"`) == false {
		t.Errorf("Expected test, got %s", out.String())
	}
	if strings.Contains(out.String(), `"ResponseBody":"foo"`) == false {
		t.Errorf("Expected test, got %s", out.String())
	}
	if strings.Contains(out.String(), `"ResponseBody":"bar"`) == false {
		t.Errorf("Expected test, got %s", out.String())
	}
	if strings.Contains(out.String(), `"Success":true`) == false {
		t.Errorf("Expected test, got %s", out.String())
	}
	if strings.Contains(out.String(), `"StatusCode":200`) == false {
		t.Errorf("Expected test, got %s", out.String())
	}
	if strings.Contains(out.String(), `"StatusCode":23`) == false {
		t.Errorf("Expected test, got %s", out.String())
	}
	if strings.Contains(out.String(), `"StatusCode":42`) == false {
		t.Errorf("Expected test, got %s", out.String())
	}
}

func TestJsonSenderFile(t *testing.T) {
	// section 1 write to file
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
	message := salvoJson.Message{
		Results: []runner.Result{resultOne, resultTwo, resultThree},
		Fails:   0,
	}
	path := "./"
	info, err := os.Stat(path)
	if err != nil {
		t.Errorf("Error getting file info: %s", err)
	}
	path, err = filepath.Abs(info.Name())
	if err != nil {
		t.Errorf("Error getting absolute path: %s", err)
	}
	err = salvoJson.SendFile(message, path)
	if err != nil {
		t.Errorf("Error sending to file: %s", err)
	}

	// section 2 read from file test values and delete file
	file, err := os.Open(path + "/salvo.json")
	if err != nil {
		t.Errorf("Error opening file: %s", err)
	}
	defer file.Close()

	jsonPath, err := filepath.Abs(file.Name())
	if err != nil {
		t.Errorf("Error getting absolute path: %s", err)
	}
	json, err := os.ReadFile(jsonPath)
	if err != nil {
		t.Errorf("Error reading file: %s", err)
	}
	for _, result := range message.Results {
		if strings.Contains(
			string(json),
			fmt.Sprintf(`"ResponseBody":"%s"`, result.ResponseBody),
		) == false {
			t.Errorf("Expected %s, got %s",
				fmt.Sprintf(`"ResponseBody":"%s"`, result.ResponseBody),
				string(json))
		}
		if strings.Contains(
			string(json),
      fmt.Sprintf(`"Success":%t`, result.Success),
		) == false {
			t.Errorf("Expected %s, got %s", strconv.FormatBool(result.Success), string(json))
		}
		if strings.Contains(
			string(json),
      fmt.Sprintf(`"StatusCode":%d`, result.StatusCode),
		) == false {
			t.Errorf("Expected %d, got %s", result.StatusCode, string(json))
		}
	}

	err = os.Remove(jsonPath)
	if err != nil {
	  t.Errorf("Error removing file: %s", err)
	}
}

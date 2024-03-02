package test

import (
	// "encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/JasonBoyett/salvo/local/runner"
	"github.com/JasonBoyett/salvo/public/protobuff"
	"google.golang.org/protobuf/proto"
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

// This test requires that salvo has been built and is in the stystem path
// This test only tests weather the cli returns properly formatted protobuff
// It does not test the accuracy of the protobuff
// other tests in the suite cover the accuracy of salvo's protobuff output
func TestCliOutputProtoBuff(t *testing.T) {

	opts := runner.Opts{
		Path:         "http://localhost:" + "5081" + "/TestCli",
		Time:         1,
		Users:        1,
		Timeout:      10,
		SuccessCodes: []int{200},
		Rate:         1.0,
	}
	protoOpts := *protobuff.OptsToProtobuff(&opts)
	rawReult, err := captureStdOutProtobuff(&protoOpts, t)
	if err != nil {
		t.Errorf("Error capturing stdout: %s", err)
	}
	t.Log(string(string(rawReult)))

	// result, err := UnmarshalResultProtoBuff(rawReult, t)
	// if err != nil {
	// 	t.Errorf("Error unmarshalling result: %s", err)
	// }
	//
	// resultJson, err := json.Marshal(result)
	//
	// if !strings.Contains(string(resultJson), `"results":`) {
	// 	t.Fatalf("Expected Results to be in output\n%s", resultJson)
	// }
	// if !strings.Contains(string(resultJson), `"start":`) {
	// 	t.Fatalf("Expected Start to be in output\n%s", resultJson)
	// }
	// if !strings.Contains(string(resultJson), `"end":`) {
	// 	t.Fatalf("Expected End to be in output\n%s", resultJson)
	// }
	// if !strings.Contains(string(resultJson), `"success":`) {
	// 	t.Fatalf("Expected Success to be in output\n%s", resultJson)
	// }
	// if !strings.Contains(string(resultJson), `"statusCode":`) {
	// 	t.Fatalf("Expected StatusCode to be in output\n%s", resultJson)
	// }
	// if !strings.Contains(string(resultJson), `"responseBody":`) {
	// 	t.Fatalf("Expected ResponseBody to be in output\n%s", resultJson)
	// }
	// if !strings.Contains(string(resultJson), `"fails":`) {
	// 	t.Fatalf("Expected Fails to be in output\n%s", resultJson)
	// }
}

func captureStdOutProtobuff(p *protobuff.Options, t *testing.T) ([]byte, error) {
	var wireMessage string
	martialedProto, err := proto.Marshal(p)
	for _, b := range martialedProto {
		wireMessage += fmt.Sprintf("%d ", b)
	}
	t.Log(wireMessage)
	if err != nil {
		return []byte{}, err
	}
	args := fmt.Sprintf("-opts=%v", wireMessage)
	t.Log(args)
	cmd := exec.Command("salvo", args)
	stdout, err := cmd.Output()
	if err != nil {
		return []byte{}, err
	}
	return stdout, nil
}

func UnmarshalResultProtoBuff(data []byte, t *testing.T) (runner.Message, error) {
	var protoResponse protobuff.Response
	err := proto.Unmarshal(data, &protoResponse)
	if err != nil {
		panic(string(data))
	}
	feilds := protoResponse.ProtoReflect().Descriptor().Fields()
	for i := 0; i < feilds.Len(); i++ {
		feild := feilds.Get(i)
		protoResponse.ProtoReflect().Set(feild, protoResponse.ProtoReflect().Get(feild))
		t.Log(feilds.Get(i).Name())
	}
	return runner.Message{
		Results: []runner.Result{
			protoResponse.ProtoReflect().Get(feilds.Get(0)).Interface().(runner.Result),
		},
		Fails: 0,
	}, nil
}

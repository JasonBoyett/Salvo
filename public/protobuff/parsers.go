package protobuff

import "github.com/JasonBoyett/salvo/local/runner"

func ProtobuffToMessage(res []Result) runner.Message {
	fails := 0
	var results []runner.Result
	for i := 0; i < len(res); i++ {
		result := FromProtobuff(&(res)[i])
		if result.Success {
			fails++
		}
		results = append(results, result)
	}
	return runner.Message{
		Results: results,
		Fails:   fails,
	}
}

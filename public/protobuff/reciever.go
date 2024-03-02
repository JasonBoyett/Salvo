package protobuff

import (
	"github.com/JasonBoyett/salvo/local/runner"
)

func Receive(o *Options) runner.Message {
	runnerOpts := ParseOptsProtobuff(o)
	result, fails := runner.Run(*runnerOpts)
	return runner.Message{
		Results: result,
		Fails:   fails,
	}
}

package protobuff

import (
	"github.com/JasonBoyett/salvo/local/runner"
)

type Message struct {
	Results []runner.Result `json:"results"`
	Fails   int             `json:"fails"`
}

func Receive(o *Options) Message {
	runnerOpts := ParseOptsProtobuff(o)
	result, fails := runner.Run(*runnerOpts)
	return Message{result, fails}
}

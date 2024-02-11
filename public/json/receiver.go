package json

import (
	"encoding/json"

	"github.com/JasonBoyett/salvo/local/runner"
)

type Message struct {
	Results []runner.Result `json:"results"`
	Fails   int             `json:"fails"`
}

func Receive(data []byte) (Message, error) {
	var options runner.Opts
	err := json.Unmarshal(data, &options)
	if err != nil {
		return Message{}, err
	}

	result, fails := runner.Run(options)
	return Message{result, fails}, nil
}

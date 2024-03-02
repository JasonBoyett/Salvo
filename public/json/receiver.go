package json

import (
	"encoding/json"

	"github.com/JasonBoyett/salvo/local/runner"
)

func Receive(data []byte) (runner.Message, error) {
	var options runner.Opts
	err := json.Unmarshal(data, &options)
	if err != nil {
		return runner.Message{}, err
	}

	result, fails := runner.Run(options)
	return runner.Message{result, fails}, nil
}

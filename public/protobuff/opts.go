package protobuff

import (
  "time"
  "github.com/JasonBoyett/salvo/local/runner"
)
func OptsToProtobuff(o *runner.Opts) *Options {
  successCodes := make([]int64, len(o.SuccessCodes))
  for i, code := range o.SuccessCodes {
    successCodes[i] = int64(code)
  }
  rate := o.Rate

  return &Options{
    Path:         o.Path,
    Time:         int64(o.Time.Seconds()),
    Users:        int64(o.Users),
    Timeout:      int64(o.Timeout),
    SuccessCodes: successCodes,
    Rate:         rate,
  }
}

func ParseOptsProtobuff(protoOpts *Options) *runner.Opts {
  successCodes := make([]int, len(protoOpts.SuccessCodes))
  for i, code := range protoOpts.SuccessCodes {
    successCodes[i] = int(code)
  }

  return &runner.Opts{
    Path:         protoOpts.Path,
    Time:         time.Duration(protoOpts.Time),
    Users:        int(protoOpts.Users),
    Timeout:      int(protoOpts.Timeout),
    SuccessCodes: successCodes,
    Rate:         protoOpts.Rate,
  }
}

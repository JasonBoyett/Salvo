package protobuff

import (
	"github.com/JasonBoyett/salvo/local/runner"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToProtobuff(r runner.Result) *Result {
	start := timestamppb.Timestamp{
		Seconds: r.Start.Unix(),
		Nanos:   int32(r.Start.Nanosecond()),
	}
	end := timestamppb.Timestamp{
		Seconds: r.End.Unix(),
		Nanos:   int32(r.End.Nanosecond()),
	}
	return &Result{
		Start:        &start,
		End:          &end,
		Success:      r.Success,
		StatusCode:   int64(r.StatusCode),
		ResponseBody: r.ResponseBody,
	}
}

func FromProtobuff(r *Result) runner.Result {
	start := r.GetStart().AsTime()
	end := r.GetEnd().AsTime()
	return runner.Result{
		Start:        start,
		End:          end,
		Success:      r.GetSuccess(),
		StatusCode:   int(r.GetStatusCode()),
		ResponseBody: r.GetResponseBody(),
	}
}

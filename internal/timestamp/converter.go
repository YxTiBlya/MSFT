package timestamp

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToTimestamppb(time time.Time) *timestamppb.Timestamp {
	return &timestamppb.Timestamp{Seconds: time.Unix()}
}

package timestamp

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToTimestamppb(tm time.Time) *timestamppb.Timestamp {
	converted_time := time.Date(
		tm.Year(),
		tm.Month(),
		tm.Day(),
		tm.Hour(),
		tm.Minute(),
		tm.Second(),
		tm.Nanosecond(),
		time.Local,
	)

	return &timestamppb.Timestamp{Seconds: converted_time.Unix()}
}

package timestamp

import (
	"strconv"
	"strings"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToTimestamppb(str string) *timestamppb.Timestamp {
	splitedStr := strings.Split(str, ".")
	tmstmpSecs, _ := strconv.Atoi(splitedStr[0])
	tmstmpNans, _ := strconv.Atoi(splitedStr[1])

	return &timestamppb.Timestamp{Seconds: int64(tmstmpSecs), Nanos: int32(tmstmpNans)}
}

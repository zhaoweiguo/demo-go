package time

import (
	"time"
)

func GetTimestampMs() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

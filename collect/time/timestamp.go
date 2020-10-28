package time

import (
	"time"
)

func getTimestampMs() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

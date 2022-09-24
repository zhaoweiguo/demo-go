package timestamp

import (
	"testing"
	"time"
)

func TestGetTimestampMs(t *testing.T) {
	msTimestamp := time.Now().UnixNano() / int64(time.Millisecond)
	t.Log(msTimestamp)
}

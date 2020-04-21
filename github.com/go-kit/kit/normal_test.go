package kit

import (
	"github.com/go-kit/kit/log/level"
	"os"
	"testing"

	"github.com/go-kit/kit/log"
)

func TestSimple(t *testing.T) {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "instance_id", 123)

	logger.Log("msg", "starting")
	logger.Log("question", "what is the meaning of life?", "answer", 42)
}

func TestTimestamp(t *testing.T) {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC) // 指定打印时间

	logger.Log("msg", "starting")
	logger.Log("question", "what is the meaning of life?", "answer", 42)
}

func TestTimestampAndCaller(t *testing.T) {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC) // 指定打印时间
	logger = log.With(logger, "caller", log.DefaultCaller)   // 指定打印callback

	logger.Log("msg", "starting")
	logger.Log("question", "what is the meaning of life?", "answer", 42)
}

func TestLevel(t *testing.T) {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = level.NewFilter(logger, level.AllowInfo()) // <--指定为info级别
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	// debug级别不打印
	level.Debug(logger).Log("msg", "this message is at the debug level")
	level.Info(logger).Log("msg", "this message is at the info level")
	level.Warn(logger).Log("msg", "this message is at the warn level")
	level.Error(logger).Log("msg", "this message is at the error level")

}

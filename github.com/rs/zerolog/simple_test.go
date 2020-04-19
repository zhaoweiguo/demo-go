package zerolog

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"testing"
)

/*
zerolog allows for logging at the following levels (from highest to lowest):

panic (zerolog.PanicLevel, 5)
fatal (zerolog.FatalLevel, 4)
error (zerolog.ErrorLevel, 3)
warn (zerolog.WarnLevel, 2)
info (zerolog.InfoLevel, 1)
debug (zerolog.DebugLevel, 0)
trace (zerolog.TraceLevel, -1)
*/
func TestDemo1(t *testing.T) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Print("hello world")
	log.Info().Msg("hello world")
}

func TestDemo2(t *testing.T) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Debug().
		Str("Scale", "833 cents").
		Float64("Interval", 833.09).
		Msg("Fibonacci is everywhere")
	log.Debug().
		Str("Name", "Tom").
		Send()
}

func TestDemo3(t *testing.T) {
	tests := []struct {
		name    string
		isDebug bool
	}{
		{"debug", true}, {"default", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
			debug := tt.isDebug

			// Default level for this example is info, unless debug flag is present
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
			if debug {
				zerolog.SetGlobalLevel(zerolog.DebugLevel)
			}

			log.Debug().Msg("This message appears only when log level set to Debug")
			log.Info().Msg("This message appears when log level set to Debug or Info")

			if e := log.Debug(); e.Enabled() {
				// Compute log output only if enabled.
				value := "bar"
				e.Str("foo", value).Msg("some debug message")
			}
		})
	}

}

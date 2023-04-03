package strings

import (
	"github.com/bmizerany/assert"
	"strings"
	"testing"
)

func TestTrimSpace(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"demo1", "aaaa", "aaaa"},
		{"demo2_user", "    aaaa", "aaaa"},
		{"demo2_user", "\t\n  aaaa", "aaaa"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := strings.TrimSpace(tt.input)
			assert.Equal(t, output, tt.want)
		})
	}
}

package strings

import (
	"strings"
	"testing"

	"github.com/bmizerany/assert"
)

func TestSplit(t *testing.T) {

	tests := []struct {
		input string
		sep   string
		want  string
	}{
		{"aaa", ":", "aaa"},
		{"aaa:bbb:ccc", ":", "ccc"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			ins := strings.Split(tt.input, tt.sep)
			assert.Equal(t, ins[len(ins)-1], tt.want)
		})
	}
}

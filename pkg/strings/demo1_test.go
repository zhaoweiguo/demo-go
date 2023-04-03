package strings

import (
	"fmt"
	"github.com/bmizerany/assert"
	"strings"
	"testing"
)

func TestNormal(t *testing.T) {
	var a string
	a = "a/c/b"
	b := strings.Split(a, "/")      // b
	c := strings.HasPrefix(a, "a/") // true
	fmt.Println(b, c)
}

func TestReplace(t *testing.T) {
	old := "\u0000\u0000string\u0000"
	new := strings.ReplaceAll(old, "\u0000", "")
	assert.Equal(t, new, "string")
}

func TestIndex(t *testing.T) {
	type args struct {
		str    string
		subStr string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"demo1", args{"abcdef", "cd"}, 2},
		{"demo2_user", args{"abcdef", "abc"}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := strings.Index(tt.args.str, tt.args.subStr)
			assert.Equal(t, i, tt.want)
		})
	}

}

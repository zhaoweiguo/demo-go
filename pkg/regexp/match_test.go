package regexp

import (
	"regexp"
	"testing"
)

func TestMatchString(t *testing.T) {
	cases := []struct {
		pattern string
		str     string
		ismatch bool
	}{
		{`^*\.corge\.at`, "abc.corge.at", true},
		{"^/", "/abc", true},
		{"^/", "ab/c", false},
		{"^[A-Z]+$", "abc/cd", false},
		{"^[A-Z]+$", "ABCD", true},
	}

	for _, c := range cases {
		isMatch, err := regexp.MatchString(c.pattern, c.str)
		if err != nil {
			t.Error("[not match]", err)
		}
		if isMatch != c.ismatch {
			t.Error("[not match]", c.pattern)
		}
	}
}

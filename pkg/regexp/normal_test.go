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

func TestMustCompile(t *testing.T) {
	cases := []struct {
		pattern string
		str     string
		ismatch bool
	}{
		{`^*\.corge\.at`, "abc.corge.at", true},
	}

	for _, c := range cases {
		mustCompile := regexp.MustCompile(c.pattern)
		isMatch := mustCompile.MatchString(c.str)
		if isMatch != c.ismatch {
			t.Error("[not match]", c.pattern)
		}

	}

}

func TestCompile(t *testing.T) {

	cases := []struct {
		pattern string
		str     string
		matched string
	}{
		{`(\w|\d|_)*.jpg`, "http://img2.bdstatic.com/img/image/166314e251f95cad1c8f496ad547d3e6709c93d5197.jpg", "166314e251f95cad1c8f496ad547d3e6709c93d5197.jpg"},
		{"a(x*)b(y*)c", "axxbyyyccc   axxbyycc", "axxbyyyc"},
	}

	for _, c := range cases {
		reg, err := regexp.Compile(c.pattern)
		if err != nil {
			t.Error("[not match]", err)
		}
		name := reg.FindStringSubmatch(c.str)[0]
		if name != c.matched {
			t.Error("[not match]", c.pattern)
		}
	}

}

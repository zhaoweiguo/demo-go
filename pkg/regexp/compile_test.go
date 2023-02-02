package regexp

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

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
		assert.Equal(t, isMatch, c.ismatch)

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
		assert.NoError(t, err)
		name := reg.FindStringSubmatch(c.str)[0]
		assert.Equal(t, name, c.matched)
	}

}

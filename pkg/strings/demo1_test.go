package main

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

func TestReplace(t *testing.T)  {
	old := "\u0000\u0000string\u0000"
	new := strings.ReplaceAll(old, "\u0000", "")
	assert.Equal(t, new, "string")
}




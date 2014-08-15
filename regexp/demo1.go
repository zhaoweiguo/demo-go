package main

import (
	"fmt"
	"regexp"
)

func main() {
	// matched
	m1, e1 := regexp.MatchString("^/", "/abc")
	fmt.Printf("ismatched:%t, err:%v\n", m1, e1)
	// not matched
	m2, e2 := regexp.MatchString("^/", "ab/c")
	fmt.Printf("ismatched:%t, err:%v\n", m2, e2)
}

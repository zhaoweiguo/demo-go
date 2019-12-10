package main

import (
	"fmt"
	"regexp"
)

func main() {
	//	a := ".*\.corge\.at"
	pattern := `^*\.corge\.at`
	str := "abc.corge.at"
	isMatch, err := regexp.MatchString(pattern, str)
	fmt.Println("[MatchString]isMatch: ", isMatch, ":", err) // true : <nil>

	mustCompile := regexp.MustCompile(pattern)
	isMatch = mustCompile.MatchString(str)
	fmt.Println("[mustCompile]isMatch: ", isMatch) // true : <nil>

	imagPath := "http://img2.bdstatic.com/img/image/166314e251f95cad1c8f496ad547d3e6709c93d5197.jpg"
	reg, _ := regexp.Compile(`(\w|\d|_)*.jpg`)
	name := reg.FindStringSubmatch(imagPath)[0]
	fmt.Print("imagPath name:", name) // 166314e251f95cad1c8f496ad547d3e6709c93d5197.jpg

}

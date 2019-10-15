package main

import (
	"fmt"
	"regexp"
)

func main() {
//	a := ".*\.corge\.at"
	a := `^*\.corge\.at`
	b := "abc.corge.at"
	ok, err := regexp.MatchString(a, b)
	fmt.Println(ok, "===", err)
	fmt.Println("###################################")


	imagPath := "http://img2.bdstatic.com/img/image/166314e251f95cad1c8f496ad547d3e6709c93d5197.jpg"
	reg, _ := regexp.Compile(`(\w|\d|_)*.jpg`)
	name := reg.FindStringSubmatch(imagPath)[0]
	fmt.Print(name)    // 166314e251f95cad1c8f496ad547d3e6709c93d5197.jpg
	fmt.Println("\n###################################")


}

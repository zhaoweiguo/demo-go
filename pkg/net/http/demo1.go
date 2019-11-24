package main

import (
	"fmt"
	"net/http"
)

func main() {
	rtn1 := http.StatusText(400)
	fmt.Printf("rnt:%v\n", rtn1)
}

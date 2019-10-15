package main

import(
	"net/http"
	"fmt"
)

func main() {
	rtn1 := http.StatusText(400)
	fmt.Printf("rnt:%v\n", rtn1)
}




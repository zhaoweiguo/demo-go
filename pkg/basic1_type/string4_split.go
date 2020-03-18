package main

import (
	"fmt"
	"strings"
)

func main() {
	a := strings.Split("aa;bb;cc;dd;ee", ";")
	fmt.Printf("%d --> %s --> %s --> %s \n", len(a), a[0], a[1], a[2])

}

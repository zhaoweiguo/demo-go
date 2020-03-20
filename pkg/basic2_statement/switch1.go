package main

import (
	"fmt"
)

func main() {
	finger := 4
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3, 4, 5: //multiple expressions in case
		fmt.Println("vowel")
	default: //default case
		fmt.Println("incorrect finger number")

	}
}

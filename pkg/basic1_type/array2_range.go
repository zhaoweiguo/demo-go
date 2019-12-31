package main

import "fmt"

func main() {
	var myArray [5]int = [5]int{1, 2, 3, 4, 5}
	var mySlice []int = myArray[:3]

	// result: 1 2 3 4 5
	for _, v := range myArray { // _:索引, v:数组值
		fmt.Print(v, "")
	}
	fmt.Println()
	// result: 1 2 3
	for _, v := range mySlice {
		fmt.Print(v, "")
	}
}

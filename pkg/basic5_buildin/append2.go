package main

import "fmt"

func main() {
	mySlice := make([]int, 5, 10)

	fmt.Println(len(mySlice)) // 5
	fmt.Println(cap(mySlice)) // 10

	mySlice2 := []int{8, 9, 10}
	mySlice = append(mySlice, mySlice2...) // 相当: append(mySlice, 8, 9, 10)

}

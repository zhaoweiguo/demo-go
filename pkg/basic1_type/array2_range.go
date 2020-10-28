package main

import "log"

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	var myArray [5]int = [5]int{1, 2, 3, 4, 5}
	var mySlice []int = myArray[:3]

	// result: 1 2 3 4 5
	for _, v := range myArray { // _:索引, v:数组值
		log.Print(v, "")
	}
	log.Println()
	// result: 1 2 3
	for _, v := range mySlice {
		log.Print(v, "")
	}
}

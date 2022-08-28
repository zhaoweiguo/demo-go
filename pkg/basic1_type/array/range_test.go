package array

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}


func TestIndex(t *testing.T) {
	str := "0123456789"
	log.Println(str[:])			// 0123456789
	log.Println(str[0:])		// 0123456789
	log.Println(str[1:])		// 123456789
	log.Println(str[:1])		// 0
	log.Println(str[0:9])		// 012345678
	log.Println(str[0:10])		// 0123456789
	log.Println(str[2:5])		// 234
}

func TestRange1(t *testing.T) {
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

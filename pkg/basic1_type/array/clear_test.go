package array

import (
	"log"
	"testing"
)

func TestClear(t *testing.T)  {
	var arr = make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	log.Println(arr)
	arr = make([]int, 10)
	for i := 0; i < 5; i++ {
		arr[i] = i+10
	}
	log.Println(arr)
}



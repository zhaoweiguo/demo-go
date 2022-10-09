package array

import (
	"log"
	"testing"
)

var items = []int{1, 2, 3, 4, 5}

func TestBug(t *testing.T) {
	var all []*int
	for _, item := range items {
		all = append(all, &item)
	}
	log.Println(all) // [0xc00001c2b0 0xc00001c2b0 0xc00001c2b0 0xc00001c2b0 0xc00001c2b0]
}

func TestBug2(t *testing.T) {
	var all []*int
	for _, item := range items {
		item := item
		all = append(all, &item)
	}
	log.Println(all)
}

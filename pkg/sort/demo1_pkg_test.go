package main

import (
	"fmt"
	"log"
	"sort"
	"testing"
)

func TestReverse(t *testing.T) {
	exprs := []string{"gordon", "simon", "bland", "octopus"}
	sort.Sort(sort.Reverse(sort.StringSlice(exprs)))
	fmt.Printf("exprs:%v\n", exprs)

	// 1. type IntSlice []string
	// 2. sort.Reverse: 指定排序从高到低(默认为从低到高)
	// 3. sort.Sort: 实现排序(要求参数实现几个指定方法)
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	fmt.Printf("origin s:%v\n", s)
	sort.Reverse(sort.IntSlice(s))
	fmt.Printf("reverse s:%v\n", s)
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Printf("sorted s:%v\n", s)

}

func TestString(t *testing.T) {
	b := []string{"a", "c", "e", "f", "b", "c", "a", "g", "b", "b", "c"}
	sort.Strings(b)
	log.Println(b)
}

func TestInt(t *testing.T) {
	c := []int{1, 6, 1, 2, 4, 7, 8, 4, 3, 2, 5, 6, 6, 8}
	sort.Ints(c)
	log.Println(c)
}

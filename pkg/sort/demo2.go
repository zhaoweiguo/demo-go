package main

import(
	"sort"
	"fmt"
)

func main() {
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

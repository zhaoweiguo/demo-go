package main

import (
	"log"
)

/*
数组和切片间的转换
*/
func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}

func main() {
	demo_array_notchange()
	log.Println("==================")
	demo_array_change()
	log.Println("==================")
}

func demo_array_notchange() {
	// 切片改变数组不变
	// 本质其实是类型转换
	arr := "Gordon123"
	slice := []rune(arr) // 注意：这儿是类型转换，不是基于数组的切片
	slice[0] = '你'
	log.Printf("[%T]%v", arr, arr)
	log.Printf("[%T]%v", slice, slice)

	arr2 := arr[:3] // 这个好像还是数组
	sliceByte := []byte(arr2)
	sliceByte[0] = '3'

	sliceRune := []rune(arr2)
	sliceRune[1] = '3'
	sliceRune[2] = 23

	log.Printf("[%T]%v", arr2, arr2)
	log.Printf("[%T]%v", sliceByte, sliceByte)
	log.Printf("[%T]%v", sliceRune, sliceRune)
}

func demo_array_change() {
	// 基于数组的切片：切片改变，数组就变
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice := arr[:5]

	slice[0] = 111
	log.Printf("[%T]%v", arr, arr)
	log.Printf("[%T]%v", slice, slice)
}

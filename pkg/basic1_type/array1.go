package main

import (
	"fmt"
	"log"
)

func main() {

	// []byte, []uint8
	var val []byte = []byte(`"{\"channel\":\"buu\",\"name\":\"john\", \"msg\":\"doe\"}"`)
	for _, c := range val {
		fmt.Println(c)
	}

	// []string
	subsCodes := []string{"aaaa", "vvvvv", "dddd", "eeeee", "gfgggg"}
	for _, s := range subsCodes {
		fmt.Println(s)
	}

	// 数组新增元素
	cars := []string{"Ferrari", "Honda", "Ford"}
	cars = append(cars, "Toyota")
	log.Println(cars)

	car0 := cars[0]
	log.Println(car0)
	car1 := cars[1]
	log.Println(car1)

}

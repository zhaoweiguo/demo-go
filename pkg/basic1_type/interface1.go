package main

import (
	"log"
)

func main() {

	// int转为interface{}默认直接转
	var i int
	i = 1
	log.Println(i)
	doInterface(i)
	log.Println("------------")

	// map[string]int转为map[string]interface{}需要专门手工转
	var maps map[string]int
	maps = make(map[string]int)
	maps["a"] = 1
	maps["b"] = 2
	log.Println(maps)
	var maps2 map[string]interface{}
	maps2 = make(map[string]interface{})
	for key, value := range maps {
		maps2[key] = value
		log.Println("+")
	}
	doInterfaces(maps2)
	log.Println(maps2)


}

func doInterface(variable interface{}) {
	log.Println("++", variable)
}

func doInterfaces(variables map[string]interface{}) {
	for key, variable := range variables {
		log.Println("+++", key, variable)
	}
}

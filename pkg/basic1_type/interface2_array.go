package main

import (
	"fmt"
)

/*
interface 与 array的互转
*/

func main() {
	arr := []string{
		"1", "2", "3",
	}
	inter := ArrayToInterface(arr)
	fmt.Printf("[%T]%v\n", inter, inter) // []string可以直接当interface{}用

	inter2 := ArrayToInterfaces([]interface{}{arr})
	fmt.Printf("[%T]%v\n", inter2, inter2)

	arr2 := InterfaceToArray(inter)
	fmt.Printf("[%T]%v\n", arr2, arr2)

}

func InterfaceToArray(t interface{}) []string {
	if v, ok := t.([]string); ok {
		return v
	}
	return []string{"error"}
}

func ArrayToInterface(t interface{}) interface{} {
	return t
}
func ArrayToInterfaces(t []interface{}) interface{} {
	return t[0]
}

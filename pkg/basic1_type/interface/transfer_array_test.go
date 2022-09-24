package _interface

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	arr := []string{
		"1", "2", "3",
	}
	inter := arrayToInterface(arr)
	fmt.Printf("[%T]%v\n", inter, inter) // []string可以直接当interface{}用

	inter2 := arrayToInterfaces([]interface{}{arr})
	fmt.Printf("[%T]%v\n", inter2, inter2)

	arr2 := interfaceToArray(inter)
	fmt.Printf("[%T]%v\n", arr2, arr2)
}

func interfaceToArray(t interface{}) []string {
	if v, ok := t.([]string); ok {
		return v
	}
	return []string{"error"}
}

func arrayToInterface(t interface{}) interface{} {
	return t
}
func arrayToInterfaces(t []interface{}) interface{} {
	return t[0]
}

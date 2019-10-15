package main

import (
	"fmt"

)

type A struct {
	L string
}

var (
	ch chan interface{}
)

func main() {

	ch = make(chan interface{})
	go receive(ch)
	ch <- A{"aaa"}
}


func receive(ch chan interface{}) {
	change := <- ch
	switch ch2 := change.(type) {  // xxx.(type) 只能用于switch语句，其他参考refect
	case int:
		fmt.Printf("A.L:%v\n", ch2)
	case string:
		fmt.Println("string")
	case []uint8:
		fmt.Println("char array")
	case []string:
		fmt.Println("string array")
	case map[string]string:
		fmt.Println("map[string]string")
	default:
		fmt.Printf("unexpected type %T\n", ch2)       // %T prints whatever type t has
	}

}

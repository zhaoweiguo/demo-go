package recover

import (
	"fmt"
	"testing"
)

func TestBasic(t *testing.T)  {
	fmt.Println("before")
	defer func() {
		fmt.Println("defer")
		err := recover()
		fmt.Println("recover: ", err) // error
	}()
	panic("error")
	fmt.Println("after")
}


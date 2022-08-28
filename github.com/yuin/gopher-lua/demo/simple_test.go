package demo

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"testing"
)

var (
	L *lua.LState
)

func init()  {
	L = lua.NewState()
}

func TestNormal(t *testing.T)  {
	defer L.Close()
	if err := L.DoString(`print("hello")`); err != nil {
		panic(err)
	}
}

func TestString(t *testing.T) {
	defer L.Close()
	lv := L.Get(-1) // get the value at the top of the stack
	if str, ok := lv.(lua.LString); ok {
		// lv is LString
		fmt.Println(string(str))
	}
	fmt.Println(lv)
	if lv.Type() != lua.LTString {
		panic("string required.")
	}
}

func TestFile(t *testing.T)  {
	defer L.Close()
	if err := L.DoFile("hello.lua"); err != nil {
		panic(err)
	}
}





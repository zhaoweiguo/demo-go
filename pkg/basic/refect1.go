package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 8.6
	b := float32(2.5)

	// TypeOf
	fmt.Printf("var a %v = %v\n", reflect.TypeOf(a), a)
	fmt.Printf("var b %v = %v\n", reflect.TypeOf(b), b)

	// ValueOf
	c := "hello world"
	fmt.Printf(" c: %v = %v\n", c, reflect.ValueOf(c).String())


	type Contact struct {
		Name    string  "check:len(3, 14)"
		Id      int     "check:range(1, 999)"
	}

	d := Contact{ "gordon", 23 }
	dType := reflect.TypeOf(d)
	fmt.Printf("dType:%v\n", dType)

	if nameField, ok := dType.FieldByName("Name"); ok {
		fmt.Printf("type:%q name:%q tag:%q \npkgpath:%q anonymous:%t index:%q offset:%q \n", nameField.Type, nameField.Name, nameField.Tag, nameField.PkgPath, nameField.Anonymous, nameField.Index, nameField.Offset)
	}

	// 类型: Kind()
	// int
	e := 1
	fmt.Printf("kind e:%q\n", reflect.TypeOf(e).Kind())
	// string
	f := "string"
	fmt.Printf("kind f:%q\n", reflect.TypeOf(f).Kind())
	// bool
	g := true
	fmt.Printf("kind g:%q\n", reflect.TypeOf(g).Kind())
	//struct
	fmt.Printf("kind d:%q\n", reflect.TypeOf(d).Kind())

	// func
	h := func(a int)(string, int){ return "abc", 123 }
	hType := reflect.TypeOf(h)
	fmt.Printf("kind h:%q\n", hType.Kind())
	fmt.Println(reflect.Func)

	fmt.Println()
	// 函数参数个数: NumIn()
	fmt.Printf("Numin() h:%v\n", hType.NumIn()==1)
	// 第一个参数的类型
	fmt.Printf("Numin() h:[%T]%v\n", hType.In(0), hType.In(0))
	fmt.Printf("Numin() h:[%T]%v\n", hType.In(0).Kind(), hType.In(0).Kind())

	fmt.Println()
	// 函数返回值个数: NumOut()
	fmt.Printf("NumOut() h:%v\n", hType.NumOut()==2)
	// 第一个返回值的类型
	fmt.Printf("NumOut() h:[%T]%v\n", hType.Out(0), hType.Out(0))
	fmt.Printf("NumOut() h:%v\n", hType.Out(0).AssignableTo(reflect.TypeOf((*string)(nil)).Elem()))


}

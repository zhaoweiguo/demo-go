package main

import (
	"fmt"
	"reflect"
)

// Bird is a nice type.
type Bird struct {
	Name           string
	LifeExpectance int
}

// Fly is fly.
func (b Bird) Fly() {
	fmt.Println("I am flying...")
}
func (b Bird) FlyA(a string) {
	fmt.Println("I am flying...", a)
}

// Fly is fly.
func (b *Bird) Fly2() {
	fmt.Println("I am flying2...")
}

func main() {
	demo1_traversal()
	demo2_field()
	demo3_method()
	demo4_tag()
}

// 遍历
func demo1_traversal() {
	fmt.Println("===demo1_traversal============================")
	sparrow := &Bird{"sparrow", 3}
	s := reflect.ValueOf(sparrow).Elem()
	typeOfT := s.Type()

	fmt.Println("field num:", s.NumField())
	// 遍历字段
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("[%d]: Field.Name = %s, Type=%s -> value=%v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	fmt.Println("method num:", s.NumMethod())
	// 遍历方法
	for j := 0; j < s.NumMethod(); j++ {
		fmt.Println(s.Method(j).String())
	}
}

// 修改简单字段的值
func demo2_field() {
	fmt.Println("===demo2_field============================")
	sparrow := &Bird{"sparrow", 3}
	x := 2
	v1 := reflect.ValueOf(&x) // 非索引类型记得加指针
	v1.Elem().SetInt(100)
	fmt.Println(x)

	// 修改字段的值2
	v := reflect.ValueOf(sparrow)
	fmt.Println(v.Elem())
	can := v.Elem().CanSet()
	v.Elem().Field(1).SetInt(2)                   // 按索引找到字段并修改
	v.Elem().FieldByName("Name").SetString("abc") // 按名字找到字段并修改
	//v.Elem().Set(v.Elem())
	fmt.Println(can, sparrow)
}

//动态调用方法
func demo3_method() {
	fmt.Println("===demo3_method============================")
	sparrow := &Bird{"sparrow", 3}
	v := reflect.ValueOf(sparrow)

	// 有参数方法
	method := v.MethodByName("FlyA")
	args := []reflect.Value{reflect.ValueOf("前缀")}
	fmt.Println(method.Call(args))

	// 无参数方法
	method2 := v.MethodByName("Fly")
	args2 := []reflect.Value{}
	fmt.Println(method2.Call(args2))
}

func demo4_tag() {
	fmt.Println("===demo4_tag============================")
	var u User
	t := reflect.TypeOf(u)

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		fmt.Println(sf.Tag.Get("json"), ",", sf.Tag.Get("bson"))
	}
}

type User struct {
	Name string `json:"name" bson:"b_name"`
	Age  int    `json:"age" bson:"b_age"`
}

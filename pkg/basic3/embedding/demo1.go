package main

import (
	"fmt"
)

func main() {
	demo1()
	demo2()
	demo3()
}

func demo1() {
	fmt.Println("===demo1============================")
	ad := admin{user{"张三", "zhangsan@flysnow.org"}, "管理员"}
	fmt.Println("可以直接调用,名字为：", ad.name)
	fmt.Println("也可以通过内部类型调用,名字为：", ad.user.name)
	fmt.Println("但是新增加的属性只能直接调用，级别为：", ad.level)
}

// 内部类型user有一个sayHello方法，外部类型对其进行了覆盖，同名重写sayHello
func demo2() {
	fmt.Println("===demo2============================")
	ad := admin{user{"张三", "zhangsan@flysnow.org"}, "管理员"}
	ad.user.sayHello()
	ad.sayHello() // 外部类型对其进行了覆盖，同名重写sayHello
}

// 如果内部类型实现了某个接口，那么外部类型也被认为实现了这个接口
func demo3() {
	fmt.Println("===demo3============================")
	ad := admin{user{"张三", "zhangsan@flysnow.org"}, "管理员"}
	sayHello(ad.user) //使用user作为参数
	sayHello(ad)      //使用admin作为参数
}

type user struct {
	name  string
	email string
}

func (u user) sayHello() {
	fmt.Println("Hello，i am a user")
}

type admin struct {
	user
	level string
}

func (a admin) sayHello() {
	fmt.Println("Hello，i am a admin")
}

type Hello interface {
	hello()
}

func (u user) hello() {
	fmt.Println("Hello，i am a user")
}

func sayHello(h Hello) {
	h.hello()
}

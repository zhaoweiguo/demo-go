package main

/**
我们后续 Foo 多了一个可变属性，那么我们只需要多一个 WithXXX 的方法就可以了，而 NewFoo 函数不需要任何变化，调用方只要在指定这个可变属性的地方增加 WithXXX 就可以了，扩展性非常好。
*/
func main() {

}

type Foo struct {
	name string
	id   int
	age  int
	db   interface{}
}

// FooOption 代表可选参数
type FooOption func(foo *Foo)

// WithName 代表 Name 为可选参数
func WithName(name string) FooOption {
	return func(foo *Foo) {
		foo.name = name
	}
}

// WithAge 代表 age 为可选参数
func WithAge(age int) FooOption {
	return func(foo *Foo) {
		foo.age = age
	}
}

// WithDB 代表 db 为可选参数
func WithDB(db interface{}) FooOption {
	return func(foo *Foo) {
		foo.db = db
	}
}

// NewFoo 代表初始化
func NewFoo(id int, options ...FooOption) *Foo {
	foo := &Foo{
		name: "default",
		id:   id,
		age:  10,
		db:   nil,
	}
	for _, option := range options {
		option(foo)
	}
	return foo
}

package iota

import (
	"fmt"
)

// 注意：先看明白这个例子
const (
	a = 1 << iota // a == 1  (iota == 0)
	b = 1 << iota // b == 2  (iota == 1)
	c = 3         // c == 3  (iota == 2, unused)
	d = 1 << iota // d == 8  (iota == 3)
)

func ExampleAbcd() {
	fmt.Println(a, b, c, d)
	// Output:
	// 1 2 3 8
}

const (
	userKey = iota
	permKey
	repoKey
)

func ExampleKey1() {
	fmt.Println(userKey, permKey, repoKey)
	// Output:
	// 0 1 2
}

const (
	aaa = iota
	bbb
	ccc
)

func ExampleAbc() {
	fmt.Println(aaa, bbb, ccc)
	// Output:
	// 0 1 2
}

const (
	aaa2 = 1 + iota
	bbb2
	ccc2
	ddd2 = iota
)

func ExampleAbc2() {
	fmt.Println(aaa2, bbb2, ccc2, ddd2)
	// Output:
	// 1 2 3 3
}

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexWaiterShift = iota
)

func ExampleMutex() {
	fmt.Println(mutexLocked, mutexWoken, mutexWaiterShift)
	// Output:
	// 1 2 2
}

const (
	ddd = 1 << iota
	eee
	fff
	fff2
	ggg = iota
	hhh = iota
	iii
)

func ExampleDef() {
	fmt.Println(ddd, eee, fff, fff2)
	fmt.Println(ddd, ggg, hhh, iii)
	// Output:
	// 1 2 4 8
	// 1 4 5 6
}

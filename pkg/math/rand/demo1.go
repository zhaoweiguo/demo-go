package main

import (
	"fmt"
	"time"
	//	"strconv"
	"math/rand"
)

func main() {
	//	seed, _ := strconv.ParseInt(time.Now().Format("20060102150405"), 10, 0)
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 20; i++ {
		fmt.Println(rand.Intn(3))
	}

	for i := 0; i < 10; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano())) // 设定rand种子
		// 得到值为0.0到1.0间float32的值
		a := r.Float32()
		fmt.Printf("a:%v\n", a)

		// 得到10个元素的数组，每个数组值是0-9的整数
		b := r.Perm(10)
		fmt.Println(b)

		time.Sleep(time.Duration(100) * time.Microsecond)
	}

}

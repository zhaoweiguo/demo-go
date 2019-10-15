package main

import(
	"fmt"
	"time"
//	"strconv"
	"math/rand"
)

func main() {
//	seed, _ := strconv.ParseInt(time.Now().Format("20060102150405"), 10, 0)
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())

	for i:=0; i<20; i++ {
		fmt.Println(rand.Intn(3))
	}

	for true{
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		a := r.Float32()
		fmt.Printf("a:%v\n", a)
		time.Sleep(time.Duration(100) * time.Microsecond)
	}

}

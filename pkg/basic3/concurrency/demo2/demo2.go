package main

import (
	"bufio"
	"log"
	"os"
	"sync"
	"time"
)

type A struct {
	B string
}

var a *A
var mu sync.Mutex

func main() {
	//newA1 := GetA()
	//log.Println("newA1")
	//log.Println(newA1)
	//newA2 := GetA()
	//log.Println("newA2")
	//log.Println(newA2)
	for i := 0; i < 10; i++ {
		go doIt()
	}

	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func doIt() {
	GetA()
	//newA3 := GetA()
	//log.Println("newA3")
	//log.Println(newA3)
}

func GetA() *A {
	newa := GetA1()
	log.Printf("==%p", newa)
	log.Printf("--%p", a)
	return newa
}

func GetA1() *A {
	log.Println("111111111111111111")
	if a == nil {
		log.Println("222222222222")
		mu.Lock()
		if a == nil {
			log.Println("3333333333333333")
			a = getRealA()
		}
		mu.Unlock()
	}
	return a
}

func GetA2() *A {
	log.Println("111111111111111111")
	if a == nil {
		log.Println("222222222222")
		mu.Lock()
		log.Println("3333333333333333")
		a = getRealA()
		mu.Unlock()
	}
	return a
}

func GetA3() *A {
	log.Println("111111111111111111")
	if a == nil {
		log.Println("3333333333333333")
		a = getRealA()
	}
	return a
}

func GetA4() *A {
	log.Println("111111111111111111")
	a = getRealA()
	return a
}

func getRealA() *A {
	time.Sleep(time.Millisecond)
	return &A{}
}

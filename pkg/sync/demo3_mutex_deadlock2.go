package main

import (
	"log"
	"sync"
	"time"
)

// 在解锁之前加锁会导致死锁
func main() {
	var mutex sync.Mutex
	go sub(mutex)
	mutex.Lock()
	log.Println("Locked")
	time.Sleep(time.Minute)
}

func sub(mutex sync.Mutex) {
	time.Sleep(time.Second)
	mutex.Lock()
	log.Println("Locked again")
}

/*
Locked
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_SemacquireMutex(0xc0000aa00c, 0x0, 0x1)
...
*/

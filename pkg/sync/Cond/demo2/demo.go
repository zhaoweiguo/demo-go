package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var ready bool

func main() {
	c := sync.NewCond(&sync.Mutex{})

	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Int63n(15)) * time.Second)
			c.L.Lock()
			for !ready {
				log.Printf("运动员#%d 已准备就绪\n", i)
				c.Wait()
			}
			c.L.Unlock()
			log.Printf("运动员#%d 开跑\n", i)
		}(i)
	}

	time.Sleep(20 * time.Second)
	log.Println("所有运动员都准备就绪...")
	c.L.Lock()
	ready = true
	c.Broadcast()
	c.L.Unlock()
	//所有的运动员是否就绪
	log.Println("========> 比赛开始，3，2，1, ......")

	time.Sleep(time.Minute)
}

package main

import (
	"log"
	"sync"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {

	log.Println("0")
	tasks := []int{1, 2, 3, 4}
	var group sync.WaitGroup
	group.Add(len(tasks))
	log.Println("2")
	for _, t := range tasks {
		go func(t int) {
			log.Println("go:", t)
			defer group.Done()
		}(t)
		log.Println("t:", t)
		group.Wait()
	}

	log.Println("1")
	//group.Wait()
	//time.Sleep(time.Millisecond)
}

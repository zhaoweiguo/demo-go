package work

import (
	"log"
	"time"
)

var value int

func init() {
	log.SetFlags(log.Lshortfile)
}

func Run() {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for true {
		select {
		case <-ticker.C:
			value = doThings()
			log.Println(value)
		}
	}
}

func doThings() int {
	return time.Now().Second()
}

func Get() int {
	return value
}

package time

import (
	"log"
	"testing"
	"time"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func TestTicker(tt *testing.T) {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	log.Println(time.Now())
	time.Sleep(4 * time.Second)

	for {
		select {
		case <-ticker.C:
			log.Println(time.Now())
		}
	}
}

func TestTimer(tt *testing.T) {
	d := time.Duration(time.Second * 2)
	timer := time.NewTimer(d)
	defer timer.Stop()

	for {
		<-timer.C
		log.Println(time.Now())
		// need reset
		timer.Reset(time.Second * 2)
	}
}

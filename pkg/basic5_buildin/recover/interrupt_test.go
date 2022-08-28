package recover

import (
	"log"
	"testing"
	"time"
)

func TestRecover(t *testing.T)  {
	log.Println("start...")
	defer func() {
		log.Println("defer")
		err := recover()
		log.Println("defer: ", err)
	}()
	log.Println("doing...")

	time.Sleep(time.Minute)
}


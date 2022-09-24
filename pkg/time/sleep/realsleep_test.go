package sleep

import (
	"log"
	"testing"
	"time"
)

func TestRealSleep(t *testing.T) {

	log.Println(time.Now().Format(time.RFC3339))

	for i := 0; i < 100; i++ {
		time.Sleep(time.Millisecond * time.Duration(100))
	}

	log.Println(time.Now().Format(time.RFC3339))

}

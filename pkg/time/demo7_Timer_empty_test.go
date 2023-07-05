package time

import (
	"log"
	"testing"
	"time"
)

func TestEmpty(t *testing.T) {
	timer := time.NewTimer(0)
	<-timer.C
	log.Println("===")
	timer.Reset(3 * time.Second)
	<-timer.C
	log.Println("===")
}

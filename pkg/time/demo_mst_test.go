package time

import (
	"log"
	"testing"
	"time"
)

func TestMST(t *testing.T) {

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	date := time.Now()
	mst := date.Format("MST") //CST
	log.Printf("mst:%s, date:%s\n", mst, date)

	if date.Format("MST") == "UTC" {
		date = date.Add(time.Hour * 8)
	}
	log.Printf("mst:%s, date:%s\n", mst, date)

}

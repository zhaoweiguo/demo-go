package zone

import (
	"log"
	"testing"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func TestZone(t *testing.T) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	log.Println(loc, err)
	d := time.Date(1987, 8, 21, 0, 0, 0, 0, loc)
	log.Println(d.Format(time.RFC3339))

	log.Println("==================================")

	loc2 := time.FixedZone("CST", 8*3600)
	log.Println(loc2)

	d2 := time.Date(1987, 8, 21, 0, 0, 0, 0, loc2)
	log.Println(d2.Format(time.RFC3339))

	log.Println("==================================")

}

func TestBirth(t *testing.T) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	log.Println(loc, err)
	birthDay := time.Date(1987, 8, 21, 0, 0, 0, 0, loc)
	log.Println(birthDay.Format(time.RFC3339))

	birth := birthDay.Unix()
	now := time.Now().Unix()

	day := (now - birth) / (24 * 60 * 60)
	log.Println(day)
}

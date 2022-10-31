package time

import (
	"fmt"
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

func TestBasic(t *testing.T) {
	date := time.Now() // 2014-07-02 17:35:39.605631353 +0800 CST
	fmt.Printf("now=%s\n", date)

	a := date.Format("Jan 2, 2006 at 3:04pm (MST)") // Jul 2, 2014 at 5:43pm (CST)
	fmt.Printf("a:%s\n", a)

	e := date.Format("2006-01-02 03:04pm")
	fmt.Printf("e:%s\n", e)

	f := date.Format("20060102_150405") // yyyyMMdd_hhmmss
	fmt.Printf("f:%s\n", f)

	g := date.Format("20060102150405") // yyyyMMddhhmmss
	log.Printf("g:%s\n", g)

	g = date.Format(time.RFC3339) // "2006-01-02T15:04:05Z07:00"
	log.Printf("RFC3339:%s\n", g)

	timestamp := time.Now().Format("2006-01-02T15:04:05Z")
	log.Println("timestamp", timestamp)

	i := date.Format("2006年01月02日")
	log.Println("i", i)
}

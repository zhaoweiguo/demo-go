package time

import (
	"log"
	"testing"
	"time"
)

func Test_parse(tt *testing.T) {

	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")
	log.Println("=>", t)

	const shortForm = "2006-Jan-02"
	t, _ = time.Parse(shortForm, "2013-Feb-03")
	log.Println("=>", t)

	t, _ = time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	log.Println("=>", t)

	t, _ = time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
	log.Println("=>", t)

	_, err := time.Parse(time.RFC3339, time.RFC3339)
	log.Println("=>", err)

	t, _ = time.Parse("2006-01-02", "2019-06-13")
	log.Println("=>", t.Format("2006-01-02"))

	startDate := "2019-06-06 14:34:00.998011694 +0800 CST"
	t, err = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", startDate)
	log.Println("=>", t)

	// 增加时区
	time.ParseInLocation("20060102", startDate, time.Local)

}

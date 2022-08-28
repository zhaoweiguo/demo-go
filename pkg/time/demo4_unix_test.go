package time

import (
	"crypto/md5"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestUnix(tt *testing.T) {

	t := time.Unix(1560257875, 0)
	log.Println("t=>", t)

	timestamp := t.Unix()
	log.Println("timestamp=>", timestamp)

	timestamp2 := time.Now().Unix()
	log.Println("timestamp2=>", timestamp2)

	t2 := time.Unix(timestamp2, 0)
	log.Println("t2=>", t2)

}

func TestUnixNano(t *testing.T) {
	now := time.Now().UnixNano() / int64(time.Millisecond)
	log.Println(now)
	strNow := strconv.FormatInt(now, 10)
	msg := "haha"
	abc := strNow + "\n" + msg
	log.Println(abc)
	log.Println(md5.Sum([]byte(abc)))

	abc2 := fmt.Sprintf("%d\n%s", now, msg)
	log.Println(abc2)
	log.Println(md5.Sum([]byte(abc2)))
}

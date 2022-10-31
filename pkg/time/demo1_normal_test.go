package time

import (
	"fmt"
	"github.com/bmizerany/assert"
	"log"
	"testing"
	"time"
)

func TestAfter(t *testing.T) {
	date1 := time.Now().Add(time.Second)
	assert.Equal(t, time.Now().After(date1), false)
	time.Sleep(time.Second*2)
	assert.Equal(t, time.Now().After(date1), true)
}

func TestAddDate(t *testing.T) {
	date := time.Now() // 2014-07-02 17:35:39.605631353 +0800 CST
	// 增加、减少操作
	b := date.Add(time.Hour * 8) // 2014-07-03 01:35:39.605631353 +0800 CST
	fmt.Printf("b:%s\n", b)

	b2 := date.AddDate(0, 0, -2) //若要获取3天前的时间，则应将-2改为-3
	fmt.Printf("b:%s\n", b2)

	c := date.UTC() // 2014-07-02 17:35:39.605631353 +0800 UTC
	fmt.Printf("c:%s\n", c)
}

func TestSetDate(t *testing.T) {
	date := time.Now() // 2014-07-02 17:35:39.605631353 +0800 CST
	// 当天0点数据
	h := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	log.Println("h=", h)
}

func TestSetDate2(t *testing.T) {
	date := time.Date(2012, 3, 4, 5, 6, 7, 0, time.UTC)
	strD := date.Format("2006年01月02日03时04分05秒")
	fmt.Printf("%v\n", strD)
}

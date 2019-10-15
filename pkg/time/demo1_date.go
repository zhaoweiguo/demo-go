package main

import (
	"log"
	"time"
	"fmt"
)

func main() {
	demo1()
	demo2()
}

func demo2() {
	fmt.Println("==========demo2 start... ")
	date := time.Date(2012, 3, 4, 5, 6, 7, 0, time.UTC)
	strD := date.Format("2006年01月02日03时04分05秒")
	fmt.Printf("[%T]%v\n", strD, strD)

	fmt.Println("==========demo2 end... \n")
}

func demo1() {
	fmt.Println("==========demo1 start... ")
	date := time.Now()     // 2014-07-02 17:35:39.605631353 +0800 CST
	fmt.Printf("now=%s\n", date)

	mst := date.Format("MST")  //CST
	fmt.Printf("mst:%s\n", mst)

	a := date.Format("Jan 2, 2006 at 3:04pm (MST)")  // Jul 2, 2014 at 5:43pm (CST)
	fmt.Printf("a:%s\n", a)

	// 增加、减少操作
	b := date.Add(time.Hour * 8)  // 2014-07-03 01:35:39.605631353 +0800 CST
	fmt.Printf("b:%s\n", b)

	b2 := date.AddDate(0, 0, -2)		//若要获取3天前的时间，则应将-2改为-3
	fmt.Printf("b:%s\n", b2)


	c := date.UTC()  // 2014-07-02 17:35:39.605631353 +0800 UTC
	fmt.Printf("c:%s\n", c)

	// 时间戳
	d := date.Unix()   // unix时间戳
	fmt.Printf("d:%s\n", d)

	e := date.Format("2006-01-02 03:04pm")
	fmt.Printf("e:[%T]%s\n", e, e)

	f := date.Format("20060102_150405")  // yyyyMMdd_hhmmss
	fmt.Printf("f:[%T]%s\n", f, f)

	g := date.Format("20060102150405")  // yyyyMMddhhmmss
	log.Println("g:[%T]%s\n", g, g)

	g = date.Format(time.RFC3339)  // yyyyMMddhhmmss
	log.Println("g:[%T]%s\n", g, g)

	timestamp := time.Now().Format("2006-01-02T15:04:05Z")
	log.Println("timestamp", timestamp)


	// 当天0点数据
	h := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	log.Println("h=", h)

	i := date.Format("2006年01月02日")
	log.Println("i", i)


	fmt.Println("==========demo1 end... \n")
}

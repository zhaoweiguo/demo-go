package main

import (
	"time"
	"fmt"
)

func main() {
	date := time.Now()     // 2014-07-02 17:35:39.605631353 +0800 CST
	fmt.Printf("now=%s\n", date)

	mst := date.Format("MST")  //CST
	fmt.Printf("mst:%s\n", mst)

	a := date.Format("Jan 2, 2006 at 3:04pm (MST)")  // Jul 2, 2014 at 5:43pm (CST)
	fmt.Printf("a:%s\n", a)

	b := date.Add(time.Hour * 8)  // 2014-07-03 01:35:39.605631353 +0800 CST
	fmt.Printf("b:%s\n", b)

	c := date.UTC()  // 2014-07-02 17:35:39.605631353 +0800 UTC
	fmt.Printf("c:%s\n", c)

	d := date.Unix()   // unix时间戳
	fmt.Printf("d:%s\n", d)

	e := date.Format("2006-01-02 03:04pm")
	fmt.Printf("e:[%T]%s\n", e, e)

	f := date.Format("20060102_150405")  // yyyyMMdd_hhmmss
	fmt.Printf("f:[%T]%s\n", f, f)
}

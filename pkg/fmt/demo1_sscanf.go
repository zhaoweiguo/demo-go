package main

import (
	"fmt"
	"log"
)

/*
	1. 把前5字符赋值给s,剩下的赋值给i
	fmt.Sscanf(" 1234567 ", "%5s%d", &s, &i)
 */
func main()  {
	var s string
	var i int
	fmt.Sscanf(" 1234567 ", "%5s%d", &s, &i)
	log.Println(s, "-", i)  // 12345 - 67

	fmt.Sscanf(" 12 34 567 ", "%5s%d", &s, &i)
	log.Println(s, "-", i)  // 12 - 34


	start := "start"
	end := "end"
	interval := "interval"
	sql := fmt.Sprintf("SELECT count(value) FROM devices2 " +
		"where time >= %s and time <= %s " +
		"group by gadget_type_sum, time(%s)", start, end, interval)
	log.Println(sql)


}

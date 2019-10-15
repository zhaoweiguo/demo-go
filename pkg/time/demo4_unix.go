package main

import (
	"log"
	"time"
)

func main()  {
	t := time.Unix(1560257875, 0)
	log.Println("t=>", t)

	timestamp := t.Unix()
	log.Println("timestamp=>", timestamp)

	timestamp2 := time.Now().Unix()
	log.Println("timestamp2=>", timestamp2)

	t2 := time.Unix(timestamp2, 0)
	log.Println("t2=>", t2)




}

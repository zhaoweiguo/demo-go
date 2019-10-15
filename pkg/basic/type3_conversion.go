package main

import (
	"log"
	"strconv"
)

func main()  {
	// Golang中string、int、int64互相转换
	//同类型之间转换，比如int64到int，直接int(int64)即可；

	string := "123"
	// string到int
	int, err:=strconv.Atoi(string)
	// string到int64
	int64, err := strconv.ParseInt(string, 10, 64)
	// int到string
	string =strconv.Itoa(int)
	// int64到string
	string =strconv.FormatInt(int64,10)

	log.Println(err)

}

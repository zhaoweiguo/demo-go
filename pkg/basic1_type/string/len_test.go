package main

import (
	"log"
	"testing"
)

func TestLen(t *testing.T) {
	str := "123456"
	log.Println(len(str)) // 6

	hex := "1122334455667788"
	log.Println(len(hex)) // 16
}

package main

import (
	"log"
	"net"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}

func main() {

	c, err := net.ListenUDP("udp4", nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// write(指定 addr 的 write)
	ra, err := net.ResolveUDPAddr("udp", "59.109.217.236:3830")
	n, err := c.WriteToUDP([]byte("222"), ra)
	if err != nil {
		log.Panic(err)
	}
	log.Println("=========================", n)
	log.Println("RemoteAddr: ", c.RemoteAddr())
	log.Println("LocalAddr: ", c.LocalAddr())
}

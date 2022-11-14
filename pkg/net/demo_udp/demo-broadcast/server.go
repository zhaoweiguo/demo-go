package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	//local address
	la, err := net.ResolveUDPAddr("udp4", "0.0.0.0:3737")
	if err != nil {
		panic(err)
	}
	//listen
	c, err := net.ListenUDP("udp4", la)
	if err != nil {
		panic(err)
	}

	//read
	bs := make([]byte, 4096)
	n, ra, err := c.ReadFromUDP(bs)
	if err != nil {
		panic(err)
	}
	fmt.Println(n, ra)

	// write
	msg := []byte(ra.String())
	log.Println("LocalAddr: ", c.LocalAddr())
	log.Println("RemoteAddr: ", c.RemoteAddr())
	log.Println(c.WriteTo(msg, ra))
	//log.Println(c.WriteToUDP(msg, ra))

}

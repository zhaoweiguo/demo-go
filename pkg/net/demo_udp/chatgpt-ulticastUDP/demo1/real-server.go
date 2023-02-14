package main

import (
	"fmt"
	"net"
	"time"
)

const (
	ip   = "239.0.0.1"
	port = "9999"
)

func main() {
	addr, _ := net.ResolveUDPAddr("udp", ip+":"+port)
	conn, _ := net.ListenUDP("udp", addr)
	defer conn.Close()

	buf := []byte("Hello, client!")

	for {
		_, err := conn.WriteToUDP(buf, addr)
		if err != nil {
			fmt.Println("WriteToUDP failed:", err)
		}
		fmt.Println(".")
		time.Sleep(time.Second)
	}
}

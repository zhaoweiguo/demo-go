package main

import (
	"fmt"
	"net"
)

const (
	ip   = "239.0.0.1"
	port = "9999"
)

func main() {
	addr, _ := net.ResolveUDPAddr("udp", ip+":"+port)
	conn, _ := net.ListenMulticastUDP("udp", nil, addr)
	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr2, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("ReadFromUDP failed:", err)
		}

		fmt.Println("Received: ", string(buf[:n]), addr2)
	}
}

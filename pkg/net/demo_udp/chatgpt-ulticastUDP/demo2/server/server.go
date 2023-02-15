package main

import (
	"fmt"
	"net"
)

const (
	serverAddr = "239.0.0.251:8888"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		fmt.Println("Failed to resolve address:", err)
		return
	}

	conn, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return
	}
	defer conn.Close()

	for {
		fmt.Println("================start........")
		b := make([]byte, 1024)
		n, src, err := conn.ReadFromUDP(b)
		if err != nil {
			fmt.Println("Failed to read:", err)
			continue
		}

		fmt.Println("Received message from", src)
		fmt.Println("Message:", string(b[:n]))

		_, err = conn.WriteToUDP([]byte("Hello from server"), src)
		if err != nil {
			fmt.Println("Failed to write:", err)
			continue
		}
		fmt.Println("================done........")
	}
}

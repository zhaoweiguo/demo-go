package main

import (
	"fmt"
	"net"
)

const (
	mdnsAddr = "239.0.0.251:8888"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", mdnsAddr)
	if err != nil {
		fmt.Println("Failed to resolve address:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Failed to dial:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte("Hello from client"))
	if err != nil {
		fmt.Println("Failed to write:", err)
		return
	}

	b := make([]byte, 1024)
	n, r, err := conn.ReadFromUDP(b)
	if err != nil {
		fmt.Println("Failed to read:", err)
		return
	}

	fmt.Println("Received message from server:", string(b[:n]))
	fmt.Println(r.IP, r.Port)
}

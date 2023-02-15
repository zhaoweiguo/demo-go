package main

import (
	"fmt"
	"net"
)

const (
	serverAddr = "239.0.0.251:8888"
	clientAddr = "239.0.0.251:9999"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", serverAddr)
	clientAddr, err := net.ResolveUDPAddr("udp", clientAddr)

	// ============================================================
	// ============================================================
	// 注意 注意 注意 注意 注意 注意 注意 注意 注意 注意 注意
	// 注意：这儿client 中不需要使用 DialUDP 函数，而是使用 ListenMulticastUDP 函数加入多播组
	// 注意 注意 注意 注意 注意 注意 注意 注意 注意 注意 注意
	// ============================================================
	// ============================================================
	//conn, err := net.DialUDP("udp", nil, addr)
	conn, err := net.ListenMulticastUDP("udp", nil, clientAddr)
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return
	}

	defer conn.Close()

	_, err = conn.WriteToUDP([]byte("Hello from client"), addr)
	if err != nil {
		fmt.Println("Failed to write:", err)
		return
	}

	b := make([]byte, 1024)
	fmt.Println("================reading........")
	n, r, err := conn.ReadFromUDP(b)
	if err != nil {
		fmt.Println("Failed to read:", err)
		return
	}

	fmt.Println("Received message from server:", string(b[:n]))
	fmt.Println("server IP:port", r.IP, r.Port)
}

package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

const (
	ip   = "239.0.0.1"
	port = "9999"
	id   = "main"
)

/**
Example:
{
	id: main
	timestamp: 1234242414
	sign: fdsafjksdfafsfsadfsf
}
*/

type Message struct {
	Id        string `json:"id"`
	Timestamp int    `json:"timestamp"`
	Sign      string `json:"sign"`
}

func main() {
	addr, _ := net.ResolveUDPAddr("udp", ip+":"+port)
	conn, _ := net.ListenUDP("udp", addr)
	defer conn.Close()

	m := Message{
		Id:        id,
		Timestamp: time.Now().Second(),
		Sign:      "1122334455",
	}
	buf, _ := json.Marshal(m)

	for {
		_, err := conn.WriteToUDP(buf, addr)
		if err != nil {
			fmt.Println("WriteToUDP failed:", err)
		}
		fmt.Println(".")
		time.Sleep(time.Second)
	}
}

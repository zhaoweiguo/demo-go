package main

import (
	"log"
	"net"
)

func main() {
	//conn, error := net.Dial("tcp", "192.168.0.1:8080")    // tcp请求
	//conn, error := net.Dial("udp", "192.168.0.1:8090")    // udp请求
	conn, error := net.Dial("ip4:icmp", "www.baidu.com") // icmp链接（使用协议名称）
	//conn, error := net.Dial("ip4:1", "10.3.3.3")          // icmp链接（使用协议编号）

	log.Println(conn, error)

}

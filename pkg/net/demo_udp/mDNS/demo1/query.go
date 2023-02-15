package main

import (
	"fmt"
	"net"
)

func main() {
	// 发送mDNS查询报文
	conn, err := net.Dial("udp4", "224.0.0.251:5353")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 设置mDNS查询报文
	query := []byte{
		// 事务ID
		0x00, 0x00,
		// 标志位：递归查询
		0x01, 0x00,
		// 问题数
		0x00, 0x01,
		// 回答数
		0x00, 0x00,
		// 权威数
		0x00, 0x00,
		// 附加数
		0x00, 0x00,
		// 查询名
		0x04, 0x5f, 0x68, 0x74, 0x74, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x00,
		// 查询类型：PTR
		0x00, 0x0c,
		// 查询类：IN
		0x00, 0x01,
	}
	conn.Write(query)

	// 接收mDNS响应报文
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}

	// 解析mDNS响应报文
	name := parseName(buf, 12)
	ip := net.IPv4(buf[n-4], buf[n-3], buf[n-2], buf[n-1])
	port := int(buf[n-6])<<8 + int(buf[n-5])

	// 打印查询结果
	fmt.Printf("服务名：%s\n", name)
	fmt.Printf("IP地址：%s\n", ip.String())
	fmt.Printf("端口号：%d\n", port)
}

func parseName(buf []byte, offset int) string {
	var name string
	for {
		if buf[offset] == 0 {
			break
		}
		if buf[offset]&0xc0 == 0xc0 {
			offset = int(buf[offset]&0x3f)<<8 + int(buf[offset+1])
		} else {
			name += string(buf[offset+1:offset+1+int(buf[offset])]) + "."
			offset += int(buf[offset]) + 1
		}
	}
	return name
}

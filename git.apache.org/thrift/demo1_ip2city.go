package main

import (
	"os"
	"net"
	"fmt"
	"github.com/zhaoweiguo/demo-go/git.apache.org/thrift/ip2city"
	"git.apache.org/thrift.git/lib/go/thrift"
)


func main() {
	protoFactory := thrift.NewTBinaryProtocolFactoryDefault()

	socket, err := thrift.NewTSocket(net.JoinHostPort("10.3.255.197", "19091"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error when new socket:", err)
		os.Exit(1)
	}

	transport := thrift.NewTFramedTransport(socket)
	defer transport.Close()

	client := ip2city.NewIpServiceClientFactory(transport, protoFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "error when open transport: ", err)
		os.Exit(1)
	}

	rtn, err := client.GetIpString("1.2.3.4")
	fmt.Println(rtn)


}


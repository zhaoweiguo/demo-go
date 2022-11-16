package main

import (
	"fmt"
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
	defer c.Close()

	for true {
		//read
		bs := make([]byte, 4096)
		n, ra, err := c.ReadFromUDP(bs)
		if err != nil {
			panic(err)
		}
		fmt.Println("reading start===========")
		fmt.Println(n, "[", ra, "]> message: ", string(bs))
		//fmt.Println("reading end===========")

		// write
		fmt.Println("writing start===========")
		n, err = c.WriteToUDP(bs[:n], ra)
		if err != nil {
			panic(err)
		}
		//fmt.Println("writing end===========")
	}

}

package main

import (
	"fmt"
	"github.com/BurntSushi/toml"

)

type hostsConf struct {
	Hosts []hostConf
}
type hostConf struct {
	Name string
	Locations []locationConf
//	Locations []interface{}
}
type locationConf struct {
	Name string
	Upstream string
	Path string
}

func main(){
	var hcs hostsConf
	if _, err := toml.DecodeFile("./host.toml", &hcs); err != nil {
		panic(err.Error())
	}

	fmt.Println(hcs)

	for _, hc := range hcs.Hosts {
		fmt.Println("=========")
		fmt.Println(hc.Name)
		for _, lo := range hc.Locations {
			fmt.Print("-----")
			fmt.Println(lo)
		}
	}


}

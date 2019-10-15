package main

import(
	"fmt"
	"github.com/BurntSushi/toml"
)

type struct1 struct {
	Abc []struct2
}
type struct2 struct {
	Name string
	Cdd  []struct3
}
type struct3 struct {
	T string
}

func main() {
	var config struct1
	if _, err := toml.DecodeFile("./demo2.toml", &config); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config)

}

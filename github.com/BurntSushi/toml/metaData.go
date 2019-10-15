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
	Type string
	T    string
	Middle  toml.Primitive
}

type struct3 struct {
	A string
	B string
}

type struct4 struct {
	T string
}

func main() {
	var config struct1
	md, err := toml.DecodeFile("./metaData.toml", &config)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, s2 := range config.Abc {
		if s2.Name == "aaa" {
			var c2 struct3
			fmt.Println(2)
			fmt.Println(s2.Type)
			fmt.Println(s2.Middle)
			if err := md.PrimitiveDecode(s2.Middle, &c2); err != nil {
				fmt.Println(3)
				fmt.Println(err)
				return
			}
			fmt.Println(4)
			fmt.Println(c2)
		}
	}

}


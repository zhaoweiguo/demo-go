package main

import (
	"fmt"
	"bytes"
	"encoding/json"
	"github.com/BurntSushi/toml"
)

type Abc struct {
	A string
	B int
}

func main() {

	a := Abc{"abc", 123}

	var buf bytes.Buffer
	enc := toml.NewEncoder(&buf)
	if err:= enc.Encode(a); err != nil {
		panic(err.Error())
	}
	fmt.Printf("toml encode~ a:%s\n", buf.String())

	json := json.NewEncoder(&buf)
	if err := json.Encode(a); err != nil {
		panic(err.Error())
	}
	fmt.Printf("json encode~ a:%s\n", buf.String())
}

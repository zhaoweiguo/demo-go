package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"flag"

)

var (
	path string
)

func main() {

	arguments := os.Args[1:]

	f := flag.NewFlagSet("flag1T", -1)
	f.SetOutput(ioutil.Discard)
	f.StringVar(&path, "config", "", "path to config file")
	f.Parse(arguments)

	fmt.Printf("[%T]%v\n", f, f)

}

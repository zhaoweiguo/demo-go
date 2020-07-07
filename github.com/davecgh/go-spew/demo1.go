package main

import (
	"github.com/davecgh/go-spew/spew"
)

type Demo struct {
	Id   int
	Name string
	Desc string
	Pids []string
}

func main() {
	demo := Demo{1, "Gordon", "This is my name", []string{"pic1", "pic2"}}
	spew.Dump(demo)

}

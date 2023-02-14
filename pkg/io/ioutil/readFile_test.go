package ioutil

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestReadFile(t *testing.T) {
	f, _ := ioutil.ReadFile("ignore.txt")
	log.Println(f)
}

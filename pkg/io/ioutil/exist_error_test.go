package ioutil

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestNotExist(t *testing.T) {
	_, err := ioutil.ReadFile("./not/exist/file")
	log.Println(err)
}

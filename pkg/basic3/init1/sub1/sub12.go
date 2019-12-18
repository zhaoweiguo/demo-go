package sub1

import (
	"log"
)

var b = func() bool {
	log.Println("sub1 b")
	return true
}()

func init() {
	log.Println("sub1 init12", b)
}

func Hello12() {
	log.Println("sub1 hello12")
}

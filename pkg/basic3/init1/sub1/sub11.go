package sub1

import (
	"log"

	"github.com/zhaoweiguo/demo-go/pkg/basic3/init1/sub2"
)

var a = func() bool {
	log.Println("sub1 a")
	return true
}()

func init() {
	log.Println("sub1 init")
}

func Hello() {
	log.Println("sub1 hello", a)
	sub2.Hello()
}

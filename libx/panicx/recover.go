package panicx

import (
	"fmt"
	"log"
	"runtime/debug"
)

// Recover 防止panic导致服务挂掉
func Recover() {
	p := recover()
	if p != nil {
		switch e := p.(type) {
		case error:
			log.Println("msg: ", e, "\nstack:\n", string(debug.Stack()))
		default:
			err := fmt.Errorf("%+v", p)
			log.Println("msg: ", err, "\nstack:\n", string(debug.Stack()))
		}
	}
}

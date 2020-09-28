package sub

import "log"

const subConst = 1

var subVar = 2

func init() {
	log.Println("sub1.init")
}

func SubFunc() {
	log.Println("sub1.f")
}

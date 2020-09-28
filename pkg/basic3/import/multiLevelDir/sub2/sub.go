package sub

import "log"

const subConst = 1

var subVar = 2

func init() {
	log.Println("sub2.init")
}

func SubFunc() {
	log.Println("sub2.f")
}

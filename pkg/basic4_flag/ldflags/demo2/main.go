package main

import (
	"fmt"
	"github.com/zhaoweiguo/demo-go/pkg/basic4_flag/ldflags/demo2/app/config"
)

func main() {
	fmt.Println("build.Version:\t", config.Version)
	fmt.Println("build.Time:\t", config.BuildTime)
}

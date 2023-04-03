package main

import (
	"github.com/jinzhu/configor"
	. "log"
)

// Config for config file.
var Config = struct {
	Url string
	Pwd string
}{}

func init() {
	SetFlags(Ldate | Lshortfile)
	err := configor.Load(&Config, "./github.com/go-redis/redis/v7/demo2_user/conf/config.yml") // 注意这儿是相对整个项目的目录地址
	if err != nil {
		Println(err)
	}
}

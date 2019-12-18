package main

import (
	"fmt"
	"github.com/jinzhu/configor"
)

/**
注意: 文件config.yml中的key必须都是小写
 **/
var Config = struct {
	APPName string `default:"app name"`

	DB struct {
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Port     uint   `default:"3306"`
	}

	Contacts []struct {
		Name  string
		Email string `required:"true"`
	}
}{}

func main() {
	configor.Load(&Config, "./github.com/jinzhu/configor/config.yml") // 注意这儿是相对整个项目的目录地址
	fmt.Println("config:", Config)
}

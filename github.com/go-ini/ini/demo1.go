package main

import (
	"github.com/go-ini/ini"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	cfg, err := ini.Load("./github.com/go-ini/ini/conf/app.ini")
	if err != nil {
		log.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	secIotlog := cfg.Section("iotlog")
	log.Println(secIotlog.Name())
	log.Println(secIotlog.Keys())
	log.Println(secIotlog.Body())

	debug := cfg.Section("iotlog").Key("DEBUG")
	log.Println(debug.Name(), debug.Value(), debug.String())
	boolDebug, _ := debug.Bool() // 直接转为bool类型
	log.Println(boolDebug)
	formatter := cfg.Section("iotlog").Key("Formatter")
	log.Println(formatter.Name(), formatter.String(), formatter.Value())

	noExist := cfg.Section("iotlog").Key("NoExist")
	log.Println(noExist, noExist.String() == "") // 没有此key返回空串
}

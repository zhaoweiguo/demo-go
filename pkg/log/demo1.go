package main

import (
	"github.com/EDDYCJY/go-gin-example/pkg/file"
	"log"
	"os"
)

var(
	F *os.File
	DefaultPrefix      = ""
	logger     *log.Logger
)

func main() {
	log.Println("zwg:start")
	log.Printf("%s\n", os.Args)


	// 指定日志打印路径
	var err error
	fileName := "issue.log"
	filePath := "/tmp/"
	F, err = file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
	logger.Println("test2")

}
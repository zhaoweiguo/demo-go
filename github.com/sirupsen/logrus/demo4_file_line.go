package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
)

func main() {
	demo1()
	demo2()
	demo3()
}

func demo1() {
	logrus.Println("==demo1=============================")
	logrus.SetReportCaller(true)
	logrus.Error("hello world")

}

func demo2() {
	log := logrus.New()
	log.Println("==demo2=============================")
	log.SetReportCaller(true)
	log.SetFormatter(&logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			//repo := "github.com/zhaoweiguo"
			//repopath := fmt.Sprintf("%s/src/%s", os.Getenv("GOPATH"), repo)
			//filename := strings.Replace(f.File, repopath, "", -1)
			return f.Function, fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})

	log.Error("hello world")

}

func demo3() {
	log := logrus.New()
	log.Println("==demo3=============================")
	log.SetReportCaller(true)
	log.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return f.Function, fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})

	log.Error("hello world")
}

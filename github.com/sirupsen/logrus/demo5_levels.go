package main

import (
	"github.com/sirupsen/logrus"
)

func main() {

	logrus.Debug("Debug日志")
	logrus.Info("Info日志")
	logrus.Warn("Warn日志")
	logrus.Error("Error日志")
	logrus.Fatal("Fatal日志") //log之后会调用os.Exit(1)
	logrus.Panic("Panic日志") //log之后会panic()
}

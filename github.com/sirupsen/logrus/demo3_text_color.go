package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.FullTimestamp = true                    // 显示完整时间
	customFormatter.TimestampFormat = "2006-01-02 15:04:05" // 时间格式
	customFormatter.ForceColors = true                      // 禁止颜色显示

	logrus.SetFormatter(customFormatter)
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)

	logrus.Debug("Debug日志")
	logrus.Info("Info日志")
	logrus.Warn("Warn日志")
	logrus.Error("Error日志")
	logrus.Fatal("Fatal日志") //log之后会调用os.Exit(1)
	logrus.Panic("Panic日志") //log之后会panic()
}

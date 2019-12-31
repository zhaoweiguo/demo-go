package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	demo12_reuse()
	demo13_pretty()
	demo14_level()
}

func demo12_reuse() {
	logrus.Println("==demo13=============================")
	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := logrus.WithFields(logrus.Fields{"common": "common", "num": 11})
	contextLogger.Warn("I'll be logged with common and other field")
}

func demo13_pretty() {
	logrus.Println("==demo12=============================")
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true, // 设定为pretty的json
	})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)
	// 打印
	logrus.WithFields(logrus.Fields{"omg": true, "number": 122}).Warn("The group's number increased tremendously!")
}

func demo14_level() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: false, // 设定为pretty的json
	})
	logrus.Println("==demo14=============================")
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	// 设定日志级别
	logrus.SetLevel(logrus.WarnLevel)
	// 不会打印
	logrus.WithFields(logrus.Fields{"animal": "walrus", "size": 10}).Info("aaaaaa💊")
	// 打印
	logrus.WithFields(logrus.Fields{"omg": true, "number": 122}).Warn("0000000💊")

	// 此条语句之后就停止了
	logrus.WithFields(logrus.Fields{"omg": true, "number": 100}).Fatal("111111💊")

	// 不会被打印
	logrus.WithFields(logrus.Fields{"omg": true, "number": 100}).Fatal("222222💊")
}

package main

import (
	"gopkg.in/gomail.v2"
	"log"
)

func main() {
	attachEmail()
}

func attachEmail() {
	m := gomail.NewMessage()
	m.SetHeader("From", "analysys@iotmail.zhaoweiguo.com.cn")
	m.SetHeader("To", "zhaowg3@zhaoweiguo.com", "langxw1@zhaoweiguo.com")
	m.SetAddressHeader("Cc", "luxl2@zhaoweiguo.com", "雪林")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>郎哲</b> and <i>Gordon</i>!")
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtpdm.aliyun.com", 25, "analysys@iotmail.zhaoweiguo.com.cn", "AQSWdefr123456")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	log.Println("ok")
}

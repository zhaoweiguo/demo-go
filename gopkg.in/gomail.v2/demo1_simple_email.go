package main

import (
	"gopkg.in/gomail.v2"
	"log"
)

func main() {
	httpEmail()
}

func httpEmail() {
	m := gomail.NewMessage()

	m.SetHeader("From", "416177937@qq.com")
	m.SetHeader("To", "zhaowg3@zhaoweiguo.com", "langxw1@zhaoweiguo.com")
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>郎哲</b> and <i>Gordon</i>!")
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.qq.com", 587, "416177937@qq.com", "vrfliwdwcyqycafj")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	log.Println("ok")
}

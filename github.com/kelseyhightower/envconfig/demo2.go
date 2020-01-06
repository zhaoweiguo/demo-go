package main

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

/*
from: github.com/drone/drone
使用:  @todo
	DRONE_USER_CREATE=username:zhaoweiguo,admin:true ./demo2
*/
func main() {
	user := Users{}
	log.Println(user)
	err := envconfig.Process("", &user)
	if err != nil {
		log.Println(err)
	}
	log.Println(user)

}

type Users struct {
	Create UserCreate `envconfig:"DRONE_USER_CREATE"`
}
type UserCreate struct {
	Username string
	Machine  bool
	Admin    bool
	Token    string
}

package main

import (
	"github.com/go-redis/redis/v7"
	"log"
)

var client *redis.Client

func init() {
	host := "127.0.0.1:6379"
	password := ""
	client = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})

	pong, err := client.Ping().Result()
	log.Println(pong, err)
}

func main() {
	log.Println("")
	key := "key"
	value := "value"
	a := client.Set(key, value, 0)
	log.Println("redis set:", a)
	b := client.Get(key)
	log.Println("redis get:", b)
}

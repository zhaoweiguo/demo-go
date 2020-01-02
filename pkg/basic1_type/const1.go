package main

import "log"

const (
	userKey = iota
	permKey
	repoKey
)

func main() {
	log.Println(userKey, permKey, repoKey) // 0 1 2
}

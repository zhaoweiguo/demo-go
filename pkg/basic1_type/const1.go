package main

import "log"

const (
	userKey int = iota
	permKey
	repoKey
)

func main() {
	log.Println(userKey, permKey, repoKey) // 0 1 2
}

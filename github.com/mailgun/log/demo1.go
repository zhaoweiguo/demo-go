package main

import(
	"fmt"
	log "github.com/mailgun/gotools-log"
)

func main() {
	log.Init([]*log.LogConfig{&log.LogConfig{Name: "console"}})

	log.Infof("11111")
	fmt.Println(2222)
}


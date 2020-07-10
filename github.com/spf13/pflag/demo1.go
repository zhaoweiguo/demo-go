package main

import (
	flag "github.com/spf13/pflag"
	"log"
)

func main() {
	var ip *int = flag.Int("flagname", 1234, "help message for flagname")

	var flagvar int
	flag.IntVar(&flagvar, "flagname2", 1234, "help message for flagname")

	log.Println(*ip)
	log.Println(flagvar)

	flag.Parse()

	log.Println("After Parse()..........")
	log.Println(*ip)
	log.Println(flagvar)

}

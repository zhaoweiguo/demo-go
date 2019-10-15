package main

import (
	"fmt"
	"os"
	"github.com/codegangsta/cli"
)

func main() {
	fmt.Println("cli started ...")

	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make an explosive entrance"
	app.Action = func(c *cli.Context) {
		println("boom! I say!")
	}

	app.Run(os.Args)

}

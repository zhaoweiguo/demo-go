package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

/*
 ./demo5_flag --name1.name2=Gordon --level1.level2=true
 ./demo5_flag -n=Gordon -l=true
*/
func main() {
	app := cli.NewApp()
	app.Name = "Just a demo"
	app.Usage = "Just for testing urfave"
	app.Copyright = "© 2019 zhaoweiguo"
	app.Authors = []cli.Author{
		{
			Name:  "zhaoweiguo",
			Email: "admin@zhaoweiguo.com",
		},
	}
	app.Action = run
	app.Version = Version
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "level1.level2, l",
			Usage:  "level mode",
			EnvVar: "LEVEL",
		},
		cli.StringFlag{
			Name:   "name1.name2, n",
			Usage:  "name",
			EnvVar: "NAME",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// Version of cli
var Version = "0.1.1202"

//  run with args
func run(c *cli.Context) {
	level := c.Bool("level1.level2")
	name := c.String("name1.name2")

	fmt.Println("===", level, name)
}

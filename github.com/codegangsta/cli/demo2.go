package main

import (
	"github.com/codegangsta/cli"
	"os"
)
func main () {

	var tasks = []string{"cook", "clean", "laundry", "eat", "sleep", "code"}
	app := cli.NewApp()
	app.Email = "abc@163.com"
	app.Name = "demo2Name"
	app.Usage = "demo2 usage"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name: "complete",
			ShortName: "c",
			Usage: "complete a task on the list",
			Action: func(c *cli.Context) {
				println("slice:", c.StringSlice("complete"), "c.Args():", c.Args())
				println("completed-- task: ", c.Args().First())
//				println("completed-- task: ", c.Args())
			},
			// go run demo2.go --generate-bash-completion a b c
			BashComplete: func(c *cli.Context) {
				println(">bash complete")
				// This will complete if no args are passed
				if len(c.Args()) > 0 {
					println("abc")
					return
				}
				for _, t := range tasks {
					println("> bash.tasks[i]", t)
//					println(t)
				}
			},
		},
		{
			Name: "ttt",
			ShortName: "t",
			Usage: "ttt usage",
			Action: func(c *cli.Context) {
				println("test ttt action")
			},
		},
	}

	app.Run(os.Args)

}

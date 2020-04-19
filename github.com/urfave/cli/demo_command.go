package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

/*
go build demo_command.go -o demo_command
 ./demo_command add taskName
 ./demo_command add taskName --port 80 --expose doit
// subs
 ./demo_command subs add subtask1
 ./demo_command subs remove subtask1
*/
func main() {
	app := cli.NewApp()

	app.Commands = commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func commands() []cli.Command {
	return []cli.Command{
		commandAdd(),
		commandSubs(),
	}
}

func commandAdd() cli.Command {
	return cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "add a task to the list",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "port",
				Usage: "The port that exposes",
			},
			cli.BoolFlag{
				Name:  "expose",
				Usage: " If true, a public, external service is created",
			},
		},
		Action: func(c *cli.Context) error {
			port := c.Int("port")
			expose := c.Bool("expose")
			fmt.Println("added task: ", c.Args().First(), ": ", port, expose)
			return nil
		},
	}
}

func commandSubs() cli.Command {
	return cli.Command{
		Name:    "subs",
		Aliases: []string{"t"},
		Usage:   "options for task templates",
		Subcommands: []cli.Command{
			{
				Name:  "add",
				Usage: "add a new template",
				Action: func(c *cli.Context) error {
					fmt.Println("new task template: ", c.Args().First())
					return nil
				},
			},
			{
				Name:  "remove",
				Usage: "remove an existing template",
				Action: func(c *cli.Context) error {
					fmt.Println("removed task template: ", c.Args().First())
					return nil
				},
			},
		},
	}
}

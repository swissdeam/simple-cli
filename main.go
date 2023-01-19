package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "hello_cli"
	app.Usage = "Print Hello World"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "Name, n",
			Value: "World",
			Usage: "Who to say hllo to.",
		},
	}
	app.Action = func(c *cli.Context) error {
		name := c.GlobalString("Name")
		fmt.Println("hello", name)
		return nil
	}
	app.Run(os.Args)
}

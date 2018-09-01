package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "vlog-cli"
	app.Usage = "connect to vlog server and do sth"
	app.Action = func(c *cli.Context) error {
		return cli.ShowAppHelp(c)
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

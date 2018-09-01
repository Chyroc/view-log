package main

import (
	"log"
	"os"

	ucli "github.com/urfave/cli"

	"github.com/Chyroc/vlog/internal/cli"
)

func main() {
	app := ucli.NewApp()
	app.Name = "vlog-cli"
	app.Usage = "connect to vlog server and do sth"
	app.Action = func(c *ucli.Context) error {
		cli.RunApp()
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

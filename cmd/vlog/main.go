package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/urfave/cli"

	"github.com/Chyroc/vlog/internal/common"
	"github.com/Chyroc/vlog/internal/server"
)

func main() {
	app := cli.NewApp()
	app.Name = "vlog"
	app.Usage = "view log on remote server"
	app.Action = func(c *cli.Context) error {
		config := c.String("config")
		if config == "" {
			return cli.ShowAppHelp(c)
		}

		if err := common.LoadConf(config); err != nil {
			return err
		}

		fmt.Println("start: " + common.Config.HTTP.Server)
		return http.ListenAndServe(":"+strconv.Itoa(common.Config.HTTP.Port), server.New())
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "default.toml",
			Usage: "vlog server config",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Chyroc/vlog/internal/cmd"
	"github.com/Chyroc/vlog/internal/common"
)

var config string

func init() {
	flag.StringVar(&config, "config", "default.toml", "config")
	flag.Parse()
}

func main() {
	if config == "" {
		fmt.Println("config is empty")
	}

	if err := common.LoadConf(config); err != nil {
		log.Println(err)
		return
	}

	fmt.Println("start: " + common.Config.HTTP.Server)
	if err := http.ListenAndServe(":"+strconv.Itoa(common.Config.HTTP.Port), cmd.NewServer()); err != nil {
		log.Fatal(err)
	}
}

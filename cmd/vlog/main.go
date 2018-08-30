package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Chyroc/vlog/internal/common"

	"github.com/Chyroc/vlog/internal/file"
)

var filename string
var config string

func init() {
	flag.StringVar(&filename, "file", "", "filename")
	flag.StringVar(&config, "config", "default.toml", "config")
	flag.Parse()
}

func main() {
	if filename == "" {
		fmt.Println("filename is empty.")
		return
	}
	if config == "" {
		fmt.Println("config is empty")
	}

	if err := common.LoadConf(config); err != nil {
		log.Println(err)
		return
	}

	if err := file.Tail(filename, func(line string) error {
		log.Println(line)
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}

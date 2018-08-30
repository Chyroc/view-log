package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Chyroc/vlog/internal/file"
)

var filename string

func init() {
	flag.StringVar(&filename, "file", "", "filename")
	flag.Parse()
}

func main() {
	if filename == "" {
		fmt.Println("filename is empty.")
		return
	}

	if err := file.Tail(filename, func(line string) error {
		fmt.Println(line)
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}

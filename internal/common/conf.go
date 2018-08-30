package common

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

var Config *configs

type configs struct {
	HTTP  *http  `json:"http"`
	Watch *watch `json:"watch"`
}

type http struct {
	Port int `json:"port"`
}

type watch struct {
	File string `json:"file"`
}

func LoadConf(confFile string) error {
	c := new(configs)
	if _, err := toml.DecodeFile(confFile, c); err != nil {
		return err
	}

	if c.Watch.File == "" {
		return fmt.Errorf("watch.file is empty, cannot continue")
	}

	if c.HTTP.Port == 0 {
		log.Printf("http.port config is empty, use default 5609\n")
		c.HTTP.Port = 5609
	}

	Config = c
	return nil
}

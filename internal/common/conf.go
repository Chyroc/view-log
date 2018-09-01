package common

import (
	"fmt"
	"log"
	"strconv"

	"github.com/BurntSushi/toml"
)

var Config *configs

type configs struct {
	HTTP  *http  `json:"http"`
	Watch *watch `json:"watch"`
}

type http struct {
	IP     string `json:"ip"`
	Port   int    `json:"port"`
	Server string `json:"server"`
}

type watch struct {
	Dir  string `json:"dir"`
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

	if c.HTTP.Server == "" {
		if c.HTTP.Port == 0 {
			log.Printf("http.port config is empty, use default 5609\n")
			c.HTTP.Port = 5609
		}
		if c.HTTP.IP == "" {
			log.Printf("http.ip config is empty, use default 127.0.0.1\n")
			c.HTTP.IP = "127.0.0.1"
		}
		c.HTTP.Server = "http://" + c.HTTP.IP + ":" + strconv.Itoa(c.HTTP.Port)
	}

	Config = c
	return nil
}

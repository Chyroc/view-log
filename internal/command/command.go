package command

import (
	"fmt"
	"strings"

	"github.com/Chyroc/vlog/internal/common"
)

var Commands map[string]Command

type Command struct {
	Client func(args []string) (string, error)
	Server func(args []string) (string, error)

	Min int
	Max int
}

func GetCommand(cmd string) *Command {
	cmd = strings.TrimSpace(cmd)

	if c, ok := Commands[cmd]; ok {
		return &c
	}
	return nil
}

func registerCommand(s string, cmd Command) {
	if Commands == nil {
		Commands = make(map[string]Command)
	}

	Commands[s] = cmd
}

func RunWithSlice(slice []string, client bool) (string, error) {
	if len(slice) == 0 {
		return "", nil
	}

	command := slice[0]
	args := slice[1:]

	cmd := GetCommand(command)
	if cmd == nil {
		return "", nil
	}

	if len(args) < cmd.Min || len(args) > cmd.Max {
		return "", fmt.Errorf("invalid args length")
	}

	var output string
	var err error
	if client {
		output, err = cmd.Client(args)
	} else {
		output, err = cmd.Server(args)
	}
	if err != nil {
		return "", err
	}

	return output, nil
}

func RunWithString(s string, client bool) (string, error) {
	slice, err := common.SplitSpacesWithQuotes(s)
	if err != nil {
		return "", err
	}

	return RunWithSlice(slice, client)
}

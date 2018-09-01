package cli

import "strings"

type Command struct {
	Cmd func(args []string) (string, error)
}

func getCommand(cmd string) *Command {
	cmd = strings.TrimSpace(cmd)

	if c, ok := alias[cmd]; ok {
		cmd = c
	}

	if c, ok := commands[cmd]; ok {
		return &c
	}
	return nil
}

var commands = map[string]Command{
	"ls": {Cmd: lsCmd},
}

var alias = map[string]string{
	"ll": "ls",
}

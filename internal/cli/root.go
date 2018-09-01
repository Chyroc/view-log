package cli

import (
	"fmt"

	"github.com/c-bata/go-prompt"
)

func RunApp() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionTitle("vlog-cli"),
	)
	p.Run()
}

func executor(s string) {
	cmd := getCommand(s)
	if cmd == nil {
		fmt.Printf("unknown command: %s\n", s)
		return
	}

	fmt.Println(s)
}

func completer(prompt.Document) []prompt.Suggest {
	return nil
}

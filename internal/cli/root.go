package cli

import (
	"fmt"

	"github.com/c-bata/go-prompt"

	"github.com/Chyroc/vlog/internal/command"
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
	out, err := command.RunWithString(s, true)
	if err != nil {
		fmt.Println(err)
		return
	}

	if out != "" {
		fmt.Println(out)
	}
}

func completer(prompt.Document) []prompt.Suggest {
	return nil
}

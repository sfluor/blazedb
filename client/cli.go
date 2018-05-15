package client

import (
	"fmt"

	"github.com/chzyer/readline"
)

// EXIT CLI code
const EXIT = "exit"

var completer = readline.NewPrefixCompleter(
	readline.PcItem("set"),
	readline.PcItem("get"),
	readline.PcItem("update"),
	readline.PcItem("delete"),
	readline.PcItem("exit"),
)

// StartCLI start the CLI loop printing a prompt line
func (c *Client) StartCLI() {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          ">>> ",
		AutoComplete:    completer,
		InterruptPrompt: "^C",
		EOFPrompt:       EXIT,
	})

	if err != nil {
		panic(err)
	}

	defer rl.Close()
	defer c.conn.Close()

	for {
		line, err := rl.Readline()

		if err != nil {
			break
		}

		if line == EXIT {
			return
		}

		data, err := c.Queryf("%s\n", line)

		fmt.Printf("%s\n", data)
	}
}

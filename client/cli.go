package client

import (
	"bufio"
	"fmt"
	"os"

	"github.com/chzyer/readline"
)

const EXIT = "exit"

var completer = readline.NewPrefixCompleter(
	readline.PcItem("set"),
	readline.PcItem("get"),
	readline.PcItem("update"),
	readline.PcItem("delete"),
	readline.PcItem("exit"),
)

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

		fmt.Fprintln(c.conn, line)

		message, err := bufio.NewReader(c.conn).ReadString('\n')

		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't read data: %v\n", err)
		}

		fmt.Print(message)
	}
}

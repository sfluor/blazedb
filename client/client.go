package client

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/chzyer/readline"
)

type Client struct {
	conn net.Conn
}

func New(addr string) *Client {
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't connect to database server at address %s: %s\n", addr, err)
		os.Exit(1)
	}

	return &Client{conn}
}

func (c *Client) Start() {
	rl, err := readline.New(">>> ")
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
		fmt.Fprintln(c.conn, line)

		message, err := bufio.NewReader(c.conn).ReadString('\n')

		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't read data: %v\n", err)
		}

		fmt.Print(message)
	}
}

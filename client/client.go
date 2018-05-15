package client

import (
	"fmt"
	"net"
	"os"
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

package client

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"reflect"

	"github.com/sfluor/blazedb/server"
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

func (c *Client) read() ([]byte, error) {
	data, err := bufio.NewReader(c.conn).ReadBytes('\n')

	if err != nil {
		return nil, fmt.Errorf("Couldn't read data: %v", err)
	}

	return bytes.TrimSpace(data), err
}

func (c *Client) assertSuccess(data []byte) error {
	if !reflect.DeepEqual(data, []byte(server.SUCCESS)) {
		return fmt.Errorf("Operation delete failed: %s", data)
	}

	return nil
}

func (c *Client) Get(key string) ([]byte, error) {
	fmt.Fprintf(c.conn, "get %s\n", key)

	return c.read()
}

func (c *Client) Set(key string, value []byte) error {
	fmt.Fprintf(c.conn, "set %s %s\n", key, value)

	data, err := c.read()

	if err != nil {
		return err
	}

	return c.assertSuccess(data)
}

func (c *Client) Update(key string, value []byte) error {
	fmt.Fprintf(c.conn, "update %s %s\n", key, value)

	data, err := c.read()

	if err != nil {
		return err
	}

	return c.assertSuccess(data)
}

func (c *Client) Delete(key string) error {
	fmt.Fprintf(c.conn, "delete %s\n", key)

	data, err := c.read()

	if err != nil {
		return err
	}

	return c.assertSuccess(data)
}

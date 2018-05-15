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

	if bytes.Contains(data, []byte("Error:")) {
		return nil, fmt.Errorf("%s", data)
	}

	return bytes.TrimSpace(data), err
}

func (c *Client) assertSuccess(data []byte, err error) error {

	if err != nil {
		return err
	}

	if !reflect.DeepEqual(data, server.SUCCESS) {
		return fmt.Errorf("Operation delete failed: %s", data)
	}

	return nil
}

func (c *Client) Queryf(format string, a ...interface{}) ([]byte, error) {
	fmt.Fprintf(c.conn, format, a...)

	return c.read()
}

func (c *Client) Get(key string) ([]byte, error) {
	return c.Queryf("get %s\n", key)
}

func (c *Client) Set(key string, value []byte) error {
	data, err := c.Queryf("set %s %s\n", key, value)

	return c.assertSuccess(data, err)
}

func (c *Client) Update(key string, value []byte) error {
	data, err := c.Queryf("update %s %s\n", key, value)

	return c.assertSuccess(data, err)
}

func (c *Client) Delete(key string) error {
	data, err := c.Queryf("delete %s\n", key)

	return c.assertSuccess(data, err)

}

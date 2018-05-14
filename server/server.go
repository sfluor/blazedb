package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	SET    = "set"
	GET    = "get"
	UPDATE = "update"
	DELETE = "delete"
)

type command struct {
	operation string
	payload   string
	client    net.Conn
}

type Server struct {
	ln       net.Listener
	db       *database
	commands chan *command
}

func New(port uint) *Server {

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't start tcp server on port %v: %v\n", err, port)
		os.Exit(1)
	}

	return &Server{ln, newDatabase(), make(chan *command)}

}

func (srv *Server) Start() {

	go srv.handleCommands()

	for {
		conn, err := srv.ln.Accept()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't accept a tcp connection: %v\n", err)
		}

		go srv.handleConnection(conn)
	}
}

func (srv *Server) handleCommands() {
	for {
		cmd := <-srv.commands
		cmd.payload = cmd.payload[:len(cmd.payload)-1]

		switch cmd.operation {
		case GET:
			data, err := srv.db.get(cmd.payload)

			if err != nil {
				cmd.client.Write([]byte(err.Error()))
			} else {
				cmd.client.Write(data)
			}

		case SET:
			chunks := strings.SplitN(cmd.payload, " ", 2)

			if len(chunks) < 2 {
				cmd.client.Write([]byte("Missing a parameter (the syntax is \"set key value\")\n"))
				continue
			}

			err := srv.db.set(chunks[0], []byte(chunks[1]))

			if err != nil {
				cmd.client.Write([]byte(err.Error()))
			} else {
				cmd.client.Write([]byte("success\n"))
			}

		case UPDATE:
			chunks := strings.SplitN(cmd.payload, " ", 2)

			if len(chunks) < 2 {
				cmd.client.Write([]byte("Missing a parameter (the syntax is \"update key value\")\n"))
				continue
			}

			err := srv.db.update(chunks[0], []byte(chunks[1]))

			if err != nil {
				cmd.client.Write([]byte(err.Error()))
			} else {
				cmd.client.Write([]byte("success\n"))
			}
		case DELETE:
			err := srv.db.delete(cmd.payload)

			if err != nil {
				cmd.client.Write([]byte(err.Error()))
			} else {
				cmd.client.Write([]byte("success\n"))
			}
		default:
			cmd.client.Write([]byte(fmt.Sprintf("Unknown command: %s", cmd.operation)))
		}
	}
}

func (srv *Server) handleConnection(conn net.Conn) {
	fmt.Printf("New connection from: %v\n", conn.RemoteAddr())

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't read data: %v\n", err)
		}
		message = message[:len(message)-1]

		chunks := strings.SplitN(message, " ", 2)
		if len(chunks) == 2 {
			srv.commands <- &command{chunks[0], chunks[1], conn}
		} else {
			conn.Write([]byte(fmt.Sprintf("Unknown command: %v, The syntax is <command> <parameters>\n", chunks[0])))
		}

	}
}

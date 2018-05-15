package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
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
	conf     *Config
	lg       *logrus.Logger
}

func New(conf *Config) *Server {

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Port))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't start tcp server on port %v: %v\n", err, conf.Port)
		os.Exit(1)
	}

	lg := logrus.New()

	file, err := os.OpenFile(conf.LogFile, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't open the log file: %s", err)
		os.Exit(1)
	}

	lg.Out = file

	if conf.Debug == 1 {
		lg.SetLevel(logrus.DebugLevel)
	}

	return &Server{ln, newDatabase(), make(chan *command, conf.MaxQueueSize), conf, lg}

}

func (srv *Server) Start() {
	fmt.Printf("Starting a blazedb server on port %v\n", srv.conf.Port)

	go srv.handleCommands()

	for {
		conn, err := srv.ln.Accept()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't accept the tcp connection: %v\n", err)
		}

		go srv.handleConnection(conn)
	}
}

func (srv *Server) handleCommands() {
	for {
		cmd := <-srv.commands

		srv.lg.WithFields(logrus.Fields{
			"operation": cmd.operation,
			"payload":   cmd.payload,
			"client":    cmd.client.RemoteAddr(),
		}).Debug("Received a command")

		switch cmd.operation {
		case GET:
			data, err := srv.db.get(cmd.payload)
			if err != nil {
				cmd.client.Write([]byte(err.Error() + "\n"))
			} else {
				cmd.client.Write(append(data, '\n'))
			}

		case SET:
			chunks := strings.SplitN(cmd.payload, " ", 2)

			if len(chunks) < 2 {
				cmd.client.Write([]byte("Missing a parameter (the syntax is \"set key value\")\n"))
				continue
			}

			err := srv.db.set(chunks[0], []byte(chunks[1]))

			if err != nil {
				cmd.client.Write([]byte(err.Error() + "\n"))
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
				cmd.client.Write([]byte(err.Error() + "\n"))
			} else {
				cmd.client.Write([]byte("success\n"))
			}
		case DELETE:
			err := srv.db.delete(cmd.payload)

			if err != nil {
				cmd.client.Write([]byte(err.Error() + "\n"))
			} else {
				cmd.client.Write([]byte("success\n"))
			}
		default:
			cmd.client.Write([]byte(fmt.Sprintf("Unknown command: %s", cmd.operation)))
		}
	}
}

func (srv *Server) handleConnection(conn net.Conn) {
	srv.lg.WithField("client", conn.RemoteAddr()).Info("New connection")

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			if err == io.EOF {
				fmt.Fprintln(os.Stderr, "Found EOF, closing the connection.")
				conn.Close()
				break
			}
			fmt.Fprintf(os.Stderr, "Couldn't read data: %v\n", err)
		}

		chunks := strings.SplitN(strings.TrimSpace(message), " ", 2)

		if len(chunks) == 2 {
			srv.commands <- &command{chunks[0], chunks[1], conn}
		} else {
			conn.Write([]byte(fmt.Sprintf("Unknown command: %v, The syntax is <command> <parameters>\n", chunks[0])))
		}

	}
}

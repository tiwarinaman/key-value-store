package eventloop

import (
	"bufio"
	"fmt"
	"net"
	"own-redis/internal/commands"
	"own-redis/internal/storage"
	"own-redis/pkg/constants"
	"strings"
)

type EventLoop struct {
	store    *storage.Storage
	commands map[string]commands.Command
}

func NewEventLoop() *EventLoop {
	return &EventLoop{
		store: storage.NewStorage(),
		commands: map[string]commands.Command{
			constants.Set: &commands.SetCommand{},
			constants.Get: &commands.GetCommand{},
		},
	}
}

func (e *EventLoop) HandleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	reader := bufio.NewReader(conn)

	for {
		_, err := conn.Write([]byte("> "))
		if err != nil {
			panic(err)
		}
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed")
			break
		}

		input = strings.TrimSpace(input)
		parts := strings.Split(input, " ")
		cmdName := strings.ToUpper(parts[0])
		args := parts[1:]

		cmd, exists := e.commands[cmdName]
		if !exists {
			_, err := conn.Write([]byte("ERROR: Unknown command\n"))
			if err != nil {
				panic(err)
			}
			continue
		}

		result := cmd.Execute(args, e.store)
		_, err = conn.Write([]byte(result + "\n"))
		if err != nil {
			panic(err)
		}
	}
}

func (e *EventLoop) Start(port string) {

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			panic(err)
		}
	}(listener)

	fmt.Printf("Server started on port %s\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go e.HandleConnection(conn)
	}
}

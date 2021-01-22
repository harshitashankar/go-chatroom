package main

import(
	"github.com/harshitashankar/go-chatroom/server"
)

func main() {
	var s server.ChatServer
	s = server.NewServer()
	s.Listen(":3333")

	s.Start()
}
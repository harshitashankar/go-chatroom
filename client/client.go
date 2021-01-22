package client

import (
	"github.com/harshitashankar/go-chatroom/protocol"
)

type messageHandler func(string)

type ChatClient interface {
	Dial(address string) error
	Start()
	Close()
	Send(command interface{}) error
	SetName(name string) error
	SendMessage(message string) error
	Incoming() chan protocol.MessageCommand
}
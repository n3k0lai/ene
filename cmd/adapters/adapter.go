package core

import (
	"time"
)

type IAdapter interface {
	Send(m Message)
	Respond(c Conversation)
	OnMessage(m Message)
}

type Adapter struct {
	Channel     string
	MsgRate     time.Duration
	Name        string
	Port        string
	PrivatePath string // oauth
	Server      string
	Typing      bool
}

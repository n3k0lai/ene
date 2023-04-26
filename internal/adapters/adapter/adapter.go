package Adapters

import (
	Conversation "github.com/n3k0lai/ene/internal/conversation"
)

type IAdapter interface {
	// Opens a connection to the Twitch.tv IRC chat server.
	//Connect()

	// Closes a connection to the Twitch.tv IRC chat server.
	//Disconnect()

	// Gets the conversations from the connected channel
	GetConvos() []*Conversation.Conversation

	Send(m Conversation.Message)
	//Respond(m core.Message, c core.Conversation)
	OnMessage(m Conversation.Message)
	GetName() string
	// Listens to chat messages and PING request from the IRC server.
	//HandleChat() error

	// Joins a specific chat channel.
	//JoinChannel()

	// Parses credentials needed for authentication.
	//ReadCredentials() error

	// Sends a message to the connected channel.
	//Say(msg string) error

	// Attempts to keep the bot connected and handling chat.
	Start()
}

type AdapterType int64

const (
	CliAdapterType       AdapterType = 0
	TwitchAdapterType    AdapterType = 1
	DiscordAdapterType   AdapterType = 2
	ChatGptAdapterType   AdapterType = 3
	ExtensionAdapterType AdapterType = 4
	TwitterAdapterType   AdapterType = 5
)

type Adapter struct {
	Typing bool
	Type   AdapterType
	Name   string
}

func NewAdapter(t AdapterType, n string) *Adapter {
	return &Adapter{
		Type: t,
		Name: n,
	}
}

func (a *Adapter) GetName() string {
	return a.Name
}

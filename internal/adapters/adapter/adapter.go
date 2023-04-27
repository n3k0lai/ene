package Adapters

import (
	Conversation "github.com/n3k0lai/ene/internal/conversation"
	Users "github.com/n3k0lai/ene/internal/users"
)

type IAdapter interface {
	// Opens a connection to the Twitch.tv IRC chat server.
	//Connect()

	// Closes a connection to the Twitch.tv IRC chat server.
	//Disconnect()

	// Gets the conversations from the connected channel
	//GetConvoStream() <-chan Conversation.Conversation

	Send(c Conversation.Conversation)
	//Respond(m core.Message, c core.Conversation)
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
	Start() AdapterStreams
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
	Typing       bool
	Type         AdapterType
	Name         string
	BotUser      Users.User
	ConvoStream  <-chan Conversation.Conversation
	OutputStream chan<- Conversation.Conversation
}
type AdapterStreams struct {
	ConvoStream  <-chan Conversation.Conversation
	OutputStream chan<- Conversation.Conversation
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

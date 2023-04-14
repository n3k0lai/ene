package Conversation

import (
	//Adapters "github.com/n3k0lai/ene/internal/adapters"
	Users "github.com/n3k0lai/ene/internal/users"
)

type Message struct {
	Text string
	User Users.User
	//Type Adapters.AdapterType
}

func NewMessage(text string, user Users.User) *Message {
	return &Message{
		Text: text,
		User: user,
		//Type: adapterType,
	}
}

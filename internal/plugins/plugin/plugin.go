package Plugin

import (
	Conversation "github.com/n3k0lai/ene/internal/conversation"
	Users "github.com/n3k0lai/ene/internal/users"
)

type IPlugin interface {
	Test(query string) bool
	Converse(c *Conversation.Conversation) *Conversation.Conversation

	Reset()
}

type Plugin struct {
	Name    string
	BotUser Users.User
}

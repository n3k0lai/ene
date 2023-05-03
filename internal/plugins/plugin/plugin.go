package Plugin

import (
	Conversation "github.com/n3k0lai/ene/internal/conversation"
	Users "github.com/n3k0lai/ene/internal/users"
)

type IPlugin interface {
	Test(c Conversation.Conversation) bool
	Converse(c Conversation.Conversation) Conversation.Conversation
	GetName() string
	Reset()
}

type Plugin struct {
	*Conversation.Conversation
	Name    string
	BotUser Users.User
}

func (p *Plugin) GetName() string {
	return p.Name
}

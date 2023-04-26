package Spam

import (
	Conversation "github.com/n3k0lai/ene/internal/conversation"
	Plugin "github.com/n3k0lai/ene/internal/plugins/plugin"
	Users "github.com/n3k0lai/ene/internal/users"
)

type Spam struct {
	*Plugin.Plugin
}

func (s *Spam) Test(query string) bool {
	return true
}

func NewSpam(bu *Users.User) *Spam {
	return &Spam{
		Plugin: &Plugin.Plugin{
			Name:    "spam",
			BotUser: *bu,
		},
	}
}

func (s *Spam) Converse(c *Conversation.Conversation) *Conversation.Conversation {
	//add a test message to the conversation
	c.OnMessage(Conversation.NewMessage("spam test", s.BotUser))
	return c
}

func (s *Spam) Reset() {

}

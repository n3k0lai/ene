package Spam

import (
	"math/rand"
	"time"

	Conversation "github.com/n3k0lai/ene/internal/conversation"
	Plugin "github.com/n3k0lai/ene/internal/plugins/plugin"
	Users "github.com/n3k0lai/ene/internal/users"
)

type Spam struct {
	*Plugin.Plugin
	EmoteStreak []string
}

func (s *Spam) Test(c Conversation.Conversation) bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) > 70
}

func NewSpam(bu *Users.User) *Spam {
	return &Spam{
		Plugin: &Plugin.Plugin{
			Name:    "spam",
			BotUser: *bu,
		},
	}
}

func (s *Spam) Converse(c Conversation.Conversation) Conversation.Conversation {
	//add a test message to the conversation
	spamtext := c.GetLatestMessage().Text
	c.OnMessage(Conversation.NewMessage(spamtext, s.BotUser))
	return c
}

func (s *Spam) Reset() {

}

package Conversation

import Users "github.com/n3k0lai/ene/internal/users"

type Conversation struct {
	Messages    []Message
	PluginUsed  string
	AdapterUsed string
}

func NewConversation(m Message, adapterUsed string) *Conversation {
	return &Conversation{
		Messages:    []Message{m},
		AdapterUsed: adapterUsed,
	}
}

func (c *Conversation) OnMessage(m *Message) *Conversation {

	c.Messages = append(c.Messages, *m)
	return c
}

func (c *Conversation) GetLatestMessage() Message {
	return c.Messages[len(c.Messages)-1]
}

func (c *Conversation) GetLatestMessageFromUser(user Users.User) *Message {
	for i := len(c.Messages) - 1; i >= 0; i-- {
		if c.Messages[i].User == user {
			return &c.Messages[i]
		}
	}
	return nil
}

func (c *Conversation) GetPluginUsed() string {
	return c.PluginUsed
}

func (c *Conversation) SetPluginUsed(plugin string) {
	c.PluginUsed = plugin
}

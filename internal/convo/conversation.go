package Convo

//import (
	//Adapters "github.com/n3k0lai/ene/internal/adapters"
	//Plugins "github.com/n3k0lai/ene/internal/plugins"
//)

type Conversation struct {
	Messages []Message
	Typing   bool
	Plugin   Plugins.IPlugin
	Adapter  Adapters.IAdapter
}

func NewConversation(m Message) *Conversation {
	return &Conversation{
		Messages: []Message{m},
	}
}

func (c *Conversation) OnMessage(m Message) {
	if c.Typing {
		c.StopAnswer()
	}

	c.Messages = append(c.Messages, m)
	c.Respond()
}

func (c *Conversation) StopAnswer() {
	c.Typing = false
	// TODO: send kill signal to adapter
}

func (c *Conversation) Respond() {
	c.Typing = true
	var res Message
	c.Adapter.Respond(res, *c)
}

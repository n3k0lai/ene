package core

type Conversation struct {
	Messages []Message
	Typing   bool
	Plugin   IPlugin
	Adapter  IAdapter
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

}

func (c *Conversation) Respond() {
	c.Typing = true
	var res Message
	c.Adapter.Respond(res, *c)
}

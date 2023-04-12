package core

type Message struct {
	Text string
	User User
}

func NewMessage(text string, user User) *Message {
	return &Message{
		Text: text,
		User: user,
	}
}

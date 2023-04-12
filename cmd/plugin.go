package core

type IPlugin interface {
	TestTrigger(query string) bool
	GetResponse(c Conversation) Message
}
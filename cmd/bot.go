package Bot

import (
	Adapters "github.com/n3k0lai/ene/internal/adapters"
	Convo "github.com/n3k0lai/ene/internal/convo"
	Plugins "github.com/n3k0lai/ene/internal/plugins"
	Users "github.com/n3k0lai/ene/internal/users"
)

type IBot interface {

	// Attempts to keep the bot connected and handling chat.
	Start()
}

type Bot struct {
	ActiveAdapters []Adapters.IAdapter
	ActivePlugins  []Plugins.IPlugin
	BotUser        Users.User
}

type BotConfig struct {
	Adapters []string
	Plugins  []string
}

func NewBot(config BotConfig) *Bot {
	return &Bot{
		ActiveAdapters: GetAdapters(config.Adapters),
		ActivePlugins:  GetPlugins(config.Plugins),
	}
}

func (b *Bot) Start() error {
	b.ActiveAdapters.Start()
}

func (b *Bot) GetConnectMessage() Convo.Message {
	return Convo.Message{
		User: b.BotUser,
		Text: "Hello, I'm a bot!",
	}

}

func (b *Bot) HandleMessage(m Convo.Message) Convo.Conversation {
	c := Convo.Conversation{
		Adapter:  m.Adapter,
		Messages: m,
	}
	b.Converse(c)
	return c
}

func (b *Bot) Converse(c Convo.Conversation) Convo.Conversation {

	// get latest message
	// test all plugins
	for i, val := range b.ActivePlugins {
		if val.Test(c) {
			c.Plugin = val
			break
		}
	}
}

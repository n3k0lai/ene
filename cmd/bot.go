package Bot

import (
	Adapters "github.com/n3k0lai/ene/internal/adapters"
	Adapter "github.com/n3k0lai/ene/internal/adapters/adapter"
	Conversation "github.com/n3k0lai/ene/internal/conversation"
	Plugins "github.com/n3k0lai/ene/internal/plugins"
	Plugin "github.com/n3k0lai/ene/internal/plugins/plugin"
	Users "github.com/n3k0lai/ene/internal/users"
)

type IBot interface {

	// Attempts to keep the bot connected and handling chat.
	Start()
}

type Bot struct {
	ActiveAdapters []Adapter.IAdapter
	ActivePlugins  []Plugin.IPlugin
	BotUser        Users.User
}

type BotConfig struct {
	Adapters []string
	Plugins  []string
}

func NewBot(config BotConfig) *Bot {
	botUser := Users.NewUser()

	return &Bot{
		ActiveAdapters: Adapters.GetAdapters(config.Adapters, *botUser),
		ActivePlugins:  Plugins.GetPlugins(config.Plugins),
	}
}

func (b *Bot) Start() error {
	// start adapters
	for _, val := range b.ActiveAdapters {
		val.Start()
	}
	return nil
}

func (b *Bot) GetConnectMessage() Conversation.Message {
	return Conversation.Message{
		User: b.BotUser,
		Text: "Hello, I'm a bot!",
	}

}

func (b *Bot) HandleMessage(m Conversation.Message) Conversation.Conversation {
	c := Conversation.Conversation{
		//Adapter:  m.Adapter,
		Messages: []Conversation.Message{m},
	}
	b.Converse(c)
	return c
}

func (b *Bot) Converse(c Conversation.Conversation) Conversation.Conversation {

	// get latest message
	latestMessage := c.Messages[len(c.Messages)-1]
	// test all plugins

	for _, plugin := range b.ActivePlugins {
		if plugin.Test(latestMessage.Text) {
			//c.Plugin = plugin
			break
		}
	}
	return c
}

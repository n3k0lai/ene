package Bot

import (
	Adapters "github.com/n3k0lai/ene/internal/adapters"
	Adapter "github.com/n3k0lai/ene/internal/adapters/adapter"
	Conversation "github.com/n3k0lai/ene/internal/conversation"
	Lib "github.com/n3k0lai/ene/internal/lib"
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
		ActivePlugins:  Plugins.GetPlugins(config.Plugins, botUser),
	}
}

func (b *Bot) Start() error {
	Lib.GetPrefix().Printfln(Lib.GetBootMessage())
	// start adapters
	for _, val := range b.ActiveAdapters {
		val.Start()
	}

	// bot loop goroutine
	//go func() {
	for {
		// get conversations from adapters
		for _, adapter := range b.ActiveAdapters {
			Lib.GetPrefix().Printfln("Getting conversations from %s", adapter.GetName())
			// get conversations from adapter
			conversations := adapter.GetConvos()
			// handle conversations

			for _, convo := range conversations {
				Lib.GetPrefix().Printfln("Handling conversation from %s", adapter.GetName())
				// handle conversations
				b.HandleConversation(convo, adapter)
			}
		}
	}
	//}()

	//select {}
}

func (b *Bot) GetConnectMessage() Conversation.Message {
	return Conversation.Message{
		User: b.BotUser,
		Text: "Hello, new adapter!",
	}

}

func (b *Bot) HandleConversation(c *Conversation.Conversation, a Adapter.IAdapter) *Conversation.Conversation {
	c = b.Converse(c)

	// send latest message from bot to adapter
	latestMessage := c.GetLatestMessage()
	if latestMessage.User != b.BotUser {
		a.Send(*latestMessage)
	}

	return c
}

func (b *Bot) Converse(c *Conversation.Conversation) *Conversation.Conversation {

	// get latest message
	latestMessage := c.Messages[len(c.Messages)-1]
	// test all plugins

	for _, plugin := range b.ActivePlugins {
		if plugin.Test(latestMessage.Text) {
			// run plugin
			c = plugin.Converse(c)
			break
		}
	}
	return c
}

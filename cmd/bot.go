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
	Lib.GetPrefix().Printfln(Lib.GetBootMessage())
	botUser := Users.NewUser("ene")

	return &Bot{
		ActiveAdapters: Adapters.GetAdapters(config.Adapters, *botUser),
		ActivePlugins:  Plugins.GetPlugins(config.Plugins, botUser),
	}
}

func (b *Bot) Start() error {
	// start adapters

	for _, adapter := range b.ActiveAdapters {
		Lib.GetPrefix().Printfln("Starting %s adapter", adapter.GetName())
		adapterStreams := adapter.Start()

		// bot loop goroutine
		go func() {
			for {
				conversation := <-adapterStreams.ConvoStream
				// handle conversation
				conversation = b.Converse(conversation)
				// send latest message from bot to adapter
				adapterStreams.OutputStream <- conversation
			}
		}()
	}

	select {}
}

func (b *Bot) Converse(c Conversation.Conversation) Conversation.Conversation {
	// test all plugins
	for _, plugin := range b.ActivePlugins {
		if plugin.Test(c) {
			// run plugin
			c.SetPluginUsed(plugin.GetName())
			return plugin.Converse(c)
		}
	}
	return c
}

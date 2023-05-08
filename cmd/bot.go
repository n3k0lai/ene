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
	BotStyle       Lib.StyleConfig
}

func NewBot(config Lib.BotConfig) *Bot {
	config.Style.GetPrefix("cli", "core").Printfln(config.Style.BootStr)
	botUser := Users.NewUser(config.Style.Name)
	config.Print()
	return &Bot{
		ActiveAdapters: Adapters.GetAdapters(config.Adapters, *botUser, config.Twitch, config.Style),
		ActivePlugins:  Plugins.GetPlugins(config.Plugins, botUser, config.Style),
		BotUser:        *botUser,
		BotStyle:       config.Style,
	}
}

func (b *Bot) Start() error {
	// start adapters

	for _, adapter := range b.ActiveAdapters {
		b.BotStyle.GetPrefix("cli", adapter.GetName()).Printfln("Starting %s adapter", adapter.GetName())
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

package Bot

import (
	CliAdapter "ene/internal/adapters/cli"
	Trivia "ene/internal/plugins/trivia"

	Adapter "github.com/n3k0lai/ene/internal/adapters/adapter"
	Conversation "github.com/n3k0lai/ene/internal/conversation"
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

func GetAvailableAdapters() []string {
	return []string{"cli", "twitch", "discord", "chatgpt", "extension", "twitter"}
}
func GetAdapters(adapterList []string, botUser Users.User) []Adapter.IAdapter {
	var adapters []Adapter.IAdapter
	for _, val := range adapterList {
		switch val {
		case "cli":
			adapters = append(adapters, CliAdapter.NewCliAdapter(botUser))
			//case "twitch":
			//	adapters = append(adapters, NewAdapter(TwitchAdapterType, val, b))
			//case "discord":
			//	adapters = append(adapters, NewAdapter(DiscordAdapterType, val, b))
			//case "chatgpt":
			//	adapters = append(adapters, NewAdapter(ChatGptAdapterType, val, b))
			//case "extension":
			//	adapters = append(adapters, NewAdapter(ExtensionAdapterType, val, b))
			//case "twitter":
			//	adapters = append(adapters, NewAdapter(TwitterAdapterType, val, b))
		}
	}
	return adapters
}

func GetAvailablePlugins() []string {
	return []string{"trivia", "tarot"}
}

func GetPlugins(pluginList []string) []Plugin.IPlugin {
	var plugins []Plugin.IPlugin
	for _, val := range pluginList {
		switch val {
		case "trivia":
			plugins = append(plugins, Trivia.NewTrivia())
			//case "tarot":
			//	plugins = append(plugins, NewTarot())
		}
	}
	return plugins
}
func NewBot(config BotConfig) *Bot {
	botUser := Users.NewUser();

	return &Bot{
		ActiveAdapters: GetAdapters(config.Adapters, *botUser),
		ActivePlugins:  GetPlugins(config.Plugins),
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

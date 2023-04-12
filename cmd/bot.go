package core

type Bot interface {

	// Attempts to keep the bot connected and handling chat.
	Start()
}

type BotConfig struct {
	ActivePlugins []IPlugin
	TwitchUser User
}

func (b *Bot) GetConnectMessage() Message {


}

func (b *Bot) Converse(c Conversation) Conversation {

	// get latest message
	// test all plugins
}
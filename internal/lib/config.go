package Lib

type BotConfig struct {

	// Adapters to load
	Adapters []string
	// Plugins to load
	Plugins []string

	// individual adapter config
	Twitch TwitchConfig
}

type TwitchChannelConfig struct {
	ChannelName string
	Plugins     []string
	OfflineOnly bool
}

type TwitchConfig struct {
	UserName   string
	ClientId   string
	OauthToken string
	Channels   []TwitchChannelConfig
}

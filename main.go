package main

import (
	Bot "github.com/n3k0lai/ene/cmd"
	Lib "github.com/n3k0lai/ene/internal/lib"
)

func main() {
	ene := Bot.NewBot(Lib.BotConfig{
		Adapters: []string{"cli"},
		Plugins:  []string{"trivia", "spam"},
		Twitch: Lib.TwitchConfig{
			UserName: "himboTTV",
			Channels: []Lib.TwitchChannelConfig{
				{
					ChannelName: "n3k0lai",
					Plugins:     []string{"trivia", "spam"},
					OfflineOnly: false,
				},
			},
		},
	})

	ene.Start()
}

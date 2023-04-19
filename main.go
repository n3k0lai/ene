package main

import (
	Bot "github.com/n3k0lai/ene/cmd"
)

func main() {
	ene := Bot.NewBot(Bot.BotConfig{
		Adapters: []string{"cli"},
		Plugins:  []string{"trivia"},
	})

	ene.Start()
}

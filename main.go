package main

import (
	Bot "github.com/n3k0lai/ene/cmd"
	Lib "github.com/n3k0lai/ene/internal/lib"
)

func main() {
	botConfig := Lib.GetConfig()

	bot := Bot.NewBot(botConfig)
	bot.Start()
}

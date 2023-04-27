package Adapters

import (
	Adapter "github.com/n3k0lai/ene/internal/adapters/adapter"
	CliAdapter "github.com/n3k0lai/ene/internal/adapters/cli"
	Users "github.com/n3k0lai/ene/internal/users"
	"github.com/pterm/pterm"
)

func GetAvailableAdapters() []string {
	return []string{"cli", "twitch", "discord", "chatgpt", "extension", "twitter"}
}
func GetAdapters(adapterList []string, botUser Users.User) []Adapter.IAdapter {
	var adapters []Adapter.IAdapter

	consoleUser := Users.NewUser("n3k0")
	for _, val := range adapterList {
		switch val {
		case "cli":
			adapters = append(adapters, CliAdapter.NewCliAdapter(botUser, *consoleUser))
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

	pterm.Info.Printf("Loaded %d adapters\n", len(adapters))
	return adapters
}

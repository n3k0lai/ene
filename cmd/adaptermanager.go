package Bot

import (
	Adapter "ene/internal/adapters/adapter"
	CliAdapter "ene/internal/adapters/cli"
)

func GetAvailableAdapters() []string {
	return []string{"cli", "twitch", "discord", "chatgpt", "extension", "twitter"}
}
func GetAdapters(adapterList []string) []Adapter.IAdapter {
	var adapters []Adapter.IAdapter
	for _, val := range adapterList {
		switch val {
		case "cli":
			adapters = append(adapters, CliAdapter.NewCliAdapter(Adapter))
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

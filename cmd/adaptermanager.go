package Bot

import (
	Adapters "ene/internal/adapters"
)

func GetAvailableAdapters() []string {
	return []string{"cli", "twitch", "discord", "chatgpt", "extension", "twitter"}
}
func GetAdapters(adapterList []string) []Adapters.IAdapter {
	var adapters []Adapters.IAdapter
	for _, val := range adapterList {
		switch val {
		case "cli":
			adapters = append(adapters, CliAdapter.NewCliAdapter(AdminUser, b))
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

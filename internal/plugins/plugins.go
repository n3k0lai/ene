package Plugins

import (
	Plugin "github.com/n3k0lai/ene/internal/plugins/plugin"
	Spam "github.com/n3k0lai/ene/internal/plugins/spam"
	Trivia "github.com/n3k0lai/ene/internal/plugins/trivia"
	Users "github.com/n3k0lai/ene/internal/users"
	//Tarot "github.com/n3k0lai/ene/internal/plugins/tarot"
)

func GetAvailablePlugins() []string {
	return []string{"spam", "trivia", "tarot"}
}

func GetPlugins(pluginList []string, botUser *Users.User) []Plugin.IPlugin {
	var plugins []Plugin.IPlugin
	for _, val := range pluginList {
		switch val {
		case "spam":
			plugins = append(plugins, Spam.NewSpam(botUser))
		case "trivia":
			plugins = append(plugins, Trivia.NewTrivia(botUser))
			//case "tarot":
			//	plugins = append(plugins, NewTarot())
		}
	}
	return plugins
}

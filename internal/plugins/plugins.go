package Plugins

import (
	Plugin "github.com/n3k0lai/ene/internal/plugins/plugin"
	Trivia "github.com/n3k0lai/ene/internal/plugins/trivia"
	//Tarot "github.com/n3k0lai/ene/internal/plugins/tarot"
)

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

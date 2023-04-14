package Bot

import (
	Plugins "ene/internal/plugins"
	Trivia "ene/internal/plugins/trivia"
)

func GetAvailablePlugins() []string {
	return []string{"trivia", "tarot"}
}

func GetPlugins(pluginList []string) []Plugins.IPlugin {
	var plugins []Plugins.IPlugin
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

package Lib

import (
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/viper"
)

type BotConfig struct {

	// Adapters to load
	Adapters []string
	// Plugins to load
	Plugins []string

	// individual adapter config
	Twitch TwitchConfig

	// Style to use
	Style StyleConfig
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

type StyleConfig struct {
	Name         string
	BootStr      string
	PrefixFg     pterm.Color
	PrefixBg     pterm.Color
	AltTextColor pterm.Color
}

func (BotConfig *BotConfig) Print() {
	BotConfig.Style.GetPrefix("cli", "config").Printfln("Loading config:")
	BotConfig.Style.GetPrefix("cli", "config").Printfln("style: %v", BotConfig.Style.Name)
	BotConfig.Style.GetPrefix("cli", "config").Printfln("  prefixfg: %v", BotConfig.Style.PrefixFg)
	BotConfig.Style.GetPrefix("cli", "config").Printfln("  prefixbg: %v", BotConfig.Style.PrefixBg)
	BotConfig.Style.GetPrefix("cli", "config").Printfln("  text color: %v", BotConfig.Style.AltTextColor)

	BotConfig.Style.GetPrefix("cli", "config").Printfln("adapters: %v", BotConfig.Adapters)
	BotConfig.Style.GetPrefix("cli", "config").Printfln("plugins: %v", BotConfig.Plugins)
	BotConfig.Style.GetPrefix("cli", "config").Printfln("  twitch.username: %v", BotConfig.Twitch.UserName)
	BotConfig.Style.GetPrefix("cli", "config").Printfln("  twitch.client_id: %v", BotConfig.Twitch.ClientId)
	BotConfig.Style.GetPrefix("cli", "config").Printfln("  twitch.oauth_token: %v", BotConfig.Twitch.OauthToken)
	for _, channel := range BotConfig.Twitch.Channels {
		BotConfig.Style.GetPrefix("cli", "config").Printfln("    twitch.channel.%v.channel_name: %v", channel.ChannelName, channel.ChannelName)
		BotConfig.Style.GetPrefix("cli", "config").Printfln("    twitch.channel.%v.plugins: %v", channel.ChannelName, channel.Plugins)
		BotConfig.Style.GetPrefix("cli", "config").Printfln("    twitch.channel.%v.offline_only: %v", channel.ChannelName, channel.OfflineOnly)
	}
}

func (style *StyleConfig) GetPrefixStyle() *pterm.Style {
	return pterm.NewStyle(style.PrefixFg, style.PrefixBg)
}

func (style *StyleConfig) GetPrimaryTextStyle() *pterm.Style {
	return pterm.NewStyle(style.AltTextColor, pterm.Bold)
}

func (style *StyleConfig) GetPrefix(adapterType string, pluginName string) *pterm.PrefixPrinter {
	// get prefix printer
	switch adapterType {
	case "cli":
		return pterm.PrefixPrinter.WithPrefix(

			pterm.PrefixPrinter{},
			pterm.Prefix{
				Text:  pterm.Sprintf("%v:%v", style.Name, pluginName),
				Style: style.GetPrefixStyle(),
			})
	case "twitch":
		return pterm.PrefixPrinter.WithPrefix(
			pterm.PrefixPrinter{},
			pterm.Prefix{
				Text:  "twitch:" + pluginName,
				Style: pterm.NewStyle(pterm.FgWhite, pterm.BgMagenta, pterm.Bold),
			})
	}
	return pterm.Error.WithPrefix(pterm.Prefix{
		Text:  "error:" + adapterType + ":" + pluginName,
		Style: pterm.NewStyle(pterm.FgWhite, pterm.BgRed, pterm.Bold),
	})
}

func GetConfig() BotConfig {
	viper.SetConfigName(".waifu")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.himbo")
	viper.AddConfigPath("$HOME/.waifu")
	viper.AddConfigPath("$HOME/.config/himbo")
	viper.AddConfigPath("$HOME/.config/waifu")
	viper.SetConfigType("toml")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		pterm.Error.Printfln("Error reading config file: %s", err)
		panic(err)
	}
	botStyle := StyleConfig{
		Name:         viper.GetString("style.name"),
		BootStr:      strings.Join(viper.GetStringSlice("style.bootstr"), "\n"),
		PrefixFg:     pterm.Color(viper.GetInt("style.prefixfg")),
		PrefixBg:     pterm.Color(viper.GetInt("style.prefixbg")),
		AltTextColor: pterm.Color(viper.GetInt("style.alttextcolor")),
	}
	return BotConfig{
		Style:    botStyle,
		Adapters: viper.GetStringSlice("adapters"),
		Plugins:  viper.GetStringSlice("plugins"),
		Twitch: TwitchConfig{
			UserName:   viper.GetString("twitch.username"),
			ClientId:   viper.GetString("twitch.clientid"),
			OauthToken: viper.GetString("twitch.oauthtoken"),
			Channels: []TwitchChannelConfig{
				{
					ChannelName: viper.GetString("twitch.channels.0.channelname"),
					Plugins:     viper.GetStringSlice("twitch.channels.0.plugins"),
					OfflineOnly: viper.GetBool("twitch.channels.0.offlineonly"),
				},
			},
		},
	}
}

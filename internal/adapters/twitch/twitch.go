package Twitch

import (
	"fmt"

	Twitch "github.com/gempir/go-twitch-irc/v4"
	Adapter "github.com/n3k0lai/ene/internal/adapters/adapter"
	Conversation "github.com/n3k0lai/ene/internal/conversation"
	Lib "github.com/n3k0lai/ene/internal/lib"
	Users "github.com/n3k0lai/ene/internal/users"
	"github.com/pterm/pterm"
)

type TwitchAdapter struct {
	*Adapter.Adapter
	Conversations []*Conversation.Conversation
	Config        Lib.TwitchConfig
	Client        *Twitch.Client
}

func (twtA *TwitchAdapter) GetPrefix(pluginName string, err bool) *pterm.PrefixPrinter {
	// get prefix printer
	if err {
		return pterm.Error.WithPrefix(pterm.Prefix{
			Text:  "error:twitch:" + pluginName,
			Style: pterm.NewStyle(pterm.FgWhite, pterm.BgRed, pterm.Bold),
		})
	}
	return pterm.PrefixPrinter.WithPrefix(
		pterm.PrefixPrinter{},
		pterm.Prefix{
			Text:  "twitch:" + pluginName,
			Style: pterm.NewStyle(pterm.FgWhite, pterm.BgMagenta, pterm.Bold),
		})
}
func NewTwitchAdapter(botUser Users.User, twitchConfig Lib.TwitchConfig) *TwitchAdapter {
	return &TwitchAdapter{
		Adapter: &Adapter.Adapter{
			Type:    Adapter.CliAdapterType,
			Typing:  false,
			Name:    "twitch",
			BotUser: botUser,
		},
		Config: twitchConfig,
	}
}

func (twtA *TwitchAdapter) Send(c Conversation.Conversation) {
	twtA.GetPrefix(c.GetPluginUsed(), false).Printfln(c.GetLatestMessage().Text)
	twtA.Client.Say("n3k0lai", c.GetLatestMessage().Text)
}

// Attempts to keep the bot connected and handling chat.
func (twtA *TwitchAdapter) Start() Adapter.AdapterStreams {
	convoStream := make(chan Conversation.Conversation)
	outputStream := make(chan Conversation.Conversation)
	twtA.GetPrefix("core", false).Printfln("twitch adapter started")

	// initialize and configure twitch client
	client := Twitch.NewClient(twtA.Config.UserName, fmt.Sprintf("oauth:%v", twtA.Config.OauthToken))

	client.OnPrivateMessage(func(twtm Twitch.PrivateMessage) {
		twtA.GetPrefix("chat", false).Printfln("@" + twtm.User.Name + ": " + twtm.Message)
		convoStream <- *Conversation.NewConversation(*Conversation.NewMessage(twtm.Message, *Users.GetTwitchUser(twtm.User)), twtA.Name)
		//convo := <-outputStream
		//if convo.GetLatestMessage().User == twtA.BotUser {
		//	twtA.Send(convo)
		//}
	})

	client.OnConnect(func() {
		twtA.GetPrefix("core", false).Printfln("connected to twitch")
	})

	// join channels
	var channelsToJoin []string
	for _, channel := range twtA.Config.Channels {
		channelsToJoin = append(channelsToJoin, channel.ChannelName)
	}
	client.Join(channelsToJoin...)
	twtA.GetPrefix("core", false).Printfln("joining channels: " + channelsToJoin[0])

	// initialize output stream
	go func() {
		for {
			convo := <-outputStream
			if convo.GetLatestMessage().User == twtA.BotUser {
				twtA.Send(convo)
			}
		}
	}()

	twtA.ConvoStream = convoStream
	twtA.OutputStream = outputStream

	// connect to twitch
	err := client.Connect()
	if err != nil {
		twtA.GetPrefix("core", true).Printfln("error connecting to twitch")
		panic(err)
	}
	twtA.Client = client

	return Adapter.AdapterStreams{
		ConvoStream:  convoStream,
		OutputStream: outputStream,
	}
}

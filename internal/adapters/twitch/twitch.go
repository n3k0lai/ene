package Twitch

import (
	Twitch "github.com/gempir/go-twitch-irc/v4"
	Adapter "github.com/n3k0lai/ene/internal/adapters/adapter"
	Conversation "github.com/n3k0lai/ene/internal/conversation"
	Lib "github.com/n3k0lai/ene/internal/lib"
	Users "github.com/n3k0lai/ene/internal/users"
)

type TwitchAdapter struct {
	*Adapter.Adapter
	Conversations []*Conversation.Conversation
	Config        Lib.TwitchConfig
	Client        *Twitch.Client
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
	Lib.GetPrefix(twtA.Name, c.GetPluginUsed()).Printfln(c.GetLatestMessage().Text)
	twtA.Client.Say("n3k0lai", c.GetLatestMessage().Text)
}

// Attempts to keep the bot connected and handling chat.
func (twtA *TwitchAdapter) Start() Adapter.AdapterStreams {
	convoStream := make(chan Conversation.Conversation)
	outputStream := make(chan Conversation.Conversation)
	Lib.GetPrefix(twtA.Name, "core").Printfln("twitch adapter started")

	// initialize and connect twitch
	client := Twitch.NewAnonymousClient() //NewClient(twtA.Config.UserName, twtA.Config.OauthToken)

	client.OnPrivateMessage(func(twtm Twitch.PrivateMessage) {
		Lib.GetPrefix(twtA.Name, "chat").Printfln("@" + twtm.User.Name + ": " + twtm.Message)
		convoStream <- *Conversation.NewConversation(*Conversation.NewMessage(twtm.Message, *Users.GetTwitchUser(twtm.User)), twtA.Name)
		//convo := <-outputStream
		//if convo.GetLatestMessage().User == twtA.BotUser {
		//	twtA.Send(convo)
		//}
	})
	client.OnConnect(func() {
		Lib.GetPrefix(twtA.Name, "core").Printfln("connected to twitch")
	})
	// join channels
	var channelsToJoin []string
	for _, channel := range twtA.Config.Channels {
		channelsToJoin = append(channelsToJoin, channel.ChannelName)
	}
	client.Join(channelsToJoin...)

	err := client.Connect()
	if err != nil {
		Lib.GetPrefix("error", "twitch").Printfln(err.Error() + " /n")

		panic(err)
	}

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
	twtA.Client = client

	return Adapter.AdapterStreams{
		ConvoStream:  convoStream,
		OutputStream: outputStream,
	}
}

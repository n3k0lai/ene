package Cli

import (
	Adapter "github.com/n3k0lai/ene/internal/adapters/adapter"
	Conversation "github.com/n3k0lai/ene/internal/conversation"
	Lib "github.com/n3k0lai/ene/internal/lib"
	Users "github.com/n3k0lai/ene/internal/users"
	"github.com/pterm/pterm"
)

type CliAdapter struct {
	*Adapter.Adapter
	ConsoleUser   Users.User
	Conversations []*Conversation.Conversation
}

func NewCliAdapter(botUser Users.User, consoleUser Users.User) *CliAdapter {
	return &CliAdapter{
		Adapter: &Adapter.Adapter{
			Type:    Adapter.CliAdapterType,
			Typing:  false,
			Name:    "cli",
			BotUser: botUser,
		},
		ConsoleUser: consoleUser,
	}
}

func (cli *CliAdapter) Send(c Conversation.Conversation) {
	Lib.GetPluginPrefix(c.GetPluginUsed()).Printfln(c.GetLatestMessage().Text)
}

// Attempts to keep the bot connected and handling chat.
func (cli *CliAdapter) Start() Adapter.AdapterStreams {
	convoStream := make(chan Conversation.Conversation)
	outputStream := make(chan Conversation.Conversation)
	pterm.Info.Printfln("cli adapter started")
	//logPanel := pterm.DefaultBox.WithTitle("logs").Sprint()
	//for i := 0; i < 100; i++ {
	//	logPanel.Write(fmt.Sprintf("Log message %d\n", i))
	//	time.Sleep(time.Millisecond * 100)
	//}
	//chatPanel := pterm.DefaultBox.WithTitle("chat").WithTitleBottomRight().WithRightPadding(0).WithBottomPadding(0).Sprint()
	//panels, _ := pterm.DefaultPanel.WithPanels(pterm.Panels{
	//	{{Data: logPanel}, {Data: chatPanel}},
	//	//{{Data: panel3}},
	//}).Srender()
	//pterm.DefaultBox.WithTitle("ene").Println(panels)

	// initialize input stream
	go func() {
		for {
			// get input from command line
			text, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("~>").Show()
			// create a conversation from the input
			convoStream <- *Conversation.NewConversation(*Conversation.NewMessage(text, cli.ConsoleUser), cli.Name)
			// loading indicator
			//spinner, _ := pterm.DefaultSpinner.Start("Loading...")
			convo := <-outputStream
			if convo.GetLatestMessage().User == cli.BotUser {
				cli.Send(convo)
			}
		}
	}()
	cli.ConvoStream = convoStream
	cli.OutputStream = outputStream

	return Adapter.AdapterStreams{
		ConvoStream:  convoStream,
		OutputStream: outputStream,
	}
}

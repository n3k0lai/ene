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
	User          Users.User
	Conversations []*Conversation.Conversation
}

func NewCliAdapter(user Users.User) *CliAdapter {
	return &CliAdapter{
		Adapter: &Adapter.Adapter{
			Type:   Adapter.CliAdapterType,
			Typing: false,
			Name:   "cli",
		},
		User: user,
	}
}

func (cli *CliAdapter) Send(m Conversation.Message) {
	Lib.GetPrefix().Printfln(m.Text)
}

func (cli *CliAdapter) GetConvos() []*Conversation.Conversation {
	return cli.Conversations
}

func (cli *CliAdapter) OnMessage(m Conversation.Message) {
	// create conversation from message
	cli.Conversations = append(cli.Conversations, Conversation.NewConversation(m))
	//print current conversations to cli
	Lib.GetPrefix().Printfln(m.Text)
}

// Attempts to keep the bot connected and handling chat.
func (cli *CliAdapter) Start() {
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
	//reader := bufio.NewReader(os.Stdin)
	for {
		// get input from command line
		text, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("~>").Show()
		// send the message to the bot
		cli.OnMessage(*Conversation.NewMessage(Lib.CleanString(text), cli.User))

	}
}

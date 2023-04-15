package Cli

import (
	"fmt"

	Adapter "github.com/n3k0lai/ene/internal/adapters/adapter"
	Conversation "github.com/n3k0lai/ene/internal/conversation"
	Users "github.com/n3k0lai/ene/internal/users"
)

type CliAdapter struct {
	*Adapter.Adapter
	User Users.User
}

func NewCliAdapter(user Users.User) *CliAdapter {
	return &CliAdapter{
		Adapter: &Adapter.Adapter{
			Type:   Adapter.CliAdapterType,
			Typing: false,
			Name:   "cli",
			//Bot:    bot,
		},
		User: user,
	}
}

func (cli *CliAdapter) Send(m Conversation.Message) {
	fmt.Println(m.Text)
}

func (cli *CliAdapter) OnMessage(m Conversation.Message) {

	// send the message to the bot
	//cli.OnMessage(*m)

	// make a new conversation
	//cli.Adapter.Bot.HandleMessage(m)

}

// Attempts to keep the bot connected and handling chat.
func (cli *CliAdapter) Start() {
	//cli.Send(cli.Adapter.Bot.GetConnectMessage())
	//panel1 := pterm.DefaultBox.Sprint("Lorem ipsum dolor sit amet,\nconsectetur adipiscing elit,\nsed do eiusmod tempor incididunt\nut labore et dolore\nmagna aliqua.")
	//panel2 := pterm.DefaultBox.WithTitle("title").Sprint("Ut enim ad minim veniam,\nquis nostrud exercitation\nullamco laboris\nnisi ut aliquip\nex ea commodo\nconsequat.")
	//panels, _ := pterm.DefaultPanel.WithPanels(pterm.Panels{
	//	{{Data: panel1}, {Data: panel2}},
	//	//{{Data: panel3}},
	//}).Srender()
	//pterm.DefaultBox.WithTitle("Lorem Ipsum").WithTitleBottomRight().WithRightPadding(0).WithBottomPadding(0).Println(panels)
	for {
		// get input
		//text := pterm.DefaultInput.WithLabel("Say something: ").WithDefaultText("Hello, I'm a bot!").WithPreviewWindow().WithPointer(">").WithPointerStyle(pterm.NewStyle(pterm.FgLightCyan)).WithHideOrder().WithHideCursor().WithRemoveWh
		// send the message to the bot
		//cli.OnMessage(*Conversation.NewMessage(Lib.CleanString(text), cli.User))

	}
}

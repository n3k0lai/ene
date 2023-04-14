package CliAdapter

import (
	"fmt"

	Bot "github.com/n3k0lai/ene/cmd"
	Adapters "github.com/n3k0lai/ene/internal/adapters"
	Convo "github.com/n3k0lai/ene/internal/convo"
	Lib "github.com/n3k0lai/ene/internal/lib"
	Users "github.com/n3k0lai/ene/internal/users"
	"github.com/pterm/pterm"
)

type CliAdapter struct {
	*Adapters.Adapter
	User Users.User
}

func NewCliAdapter(user Users.User, bot *Bot.Bot) *CliAdapter {
	return &CliAdapter{
		Adapter: &Adapters.Adapter{
			Type:   Adapters.CliAdapterType,
			Typing: false,
			Name:   "cli",
			Bot:    bot,
		},
		User: user,
	}
}

func (cli *CliAdapter) Send(m Convo.Message) {
	fmt.Println(m.Text)
}

func (cli *CliAdapter) OnMessage(m Convo.Message) {

	// send the message to the bot
	cli.OnMessage(*m)

	// make a new conversation
	cli.Adapter.Bot.HandleMessage(m)

}

// Attempts to keep the bot connected and handling chat.
func (cli *CliAdapter) Start() {
	cli.Send(cli.Adapter.Bot.GetConnectMessage())
	//panel1 := pterm.DefaultBox.Sprint("Lorem ipsum dolor sit amet,\nconsectetur adipiscing elit,\nsed do eiusmod tempor incididunt\nut labore et dolore\nmagna aliqua.")
	//panel2 := pterm.DefaultBox.WithTitle("title").Sprint("Ut enim ad minim veniam,\nquis nostrud exercitation\nullamco laboris\nnisi ut aliquip\nex ea commodo\nconsequat.")
	//panels, _ := pterm.DefaultPanel.WithPanels(pterm.Panels{
	//	{{Data: panel1}, {Data: panel2}},
	//	//{{Data: panel3}},
	//}).Srender()
	//pterm.DefaultBox.WithTitle("Lorem Ipsum").WithTitleBottomRight().WithRightPadding(0).WithBottomPadding(0).Println(panels)
	for {
		// get input
		text := pterm.DefaultInput.WithLabel("Say something: ").WithDefaultText("Hello, I'm a bot!").WithPreviewWindow().WithPointer(">").WithPointerStyle(pterm.NewStyle(pterm.FgLightCyan)).WithHideOrder().WithHideCursor().WithRemoveWh
		// send the message to the bot
		cli.OnMessage(*Convo.NewMessage(Lib.CleanString(text), cli.User))

	}
}

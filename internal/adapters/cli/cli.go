package CliAdapter

import (
	"fmt"
	"strings"

	"github.com/go-delve/delve/pkg/dwarf/reader"
	"github.com/n3k0lai/ene/cmd"
)

type CliAdapter struct {
	*core.Adapter
}

func NewCliAdapter() *CliAdapter {
	return &CliAdapter{
		Adapter: &core.Adapter{
			Type: core.CliAdapterType,
			Typing: false,
		},
	}
}

func Send(m core.Message) {
	fmt.Println(m.Text)
}

func Respond(c core.Conversation) {

}

func OnMessage(m core.Message) {

}

// Attempts to keep the bot connected and handling chat.
func Start() {
	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		//bot.testplugin
		//bot.runplugin
	}
}
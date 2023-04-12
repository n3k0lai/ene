package OpenAI

import (
	core "github.com/n3k0lai/ene/cmd"
	openai "github.com/sashabaranov/go-openai"
)

type OpenAiAdapter struct {
	*core.Adapter
	Client *openai.Client
}
func (oaa *OpenAiAdapter) Connect() {
	oaa.Client = openai.NewClient("your token")
}

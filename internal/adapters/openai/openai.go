package OpenAI

import (
	Adapter "github.com/n3k0lai/ene/internal/adapters/adapter"
	openai "github.com/sashabaranov/go-openai"
)

type OpenAiAdapter struct {
	*Adapter.Adapter
	Client *openai.Client
}

func (oaa *OpenAiAdapter) Connect() {
	oaa.Client = openai.NewClient("your token")
}

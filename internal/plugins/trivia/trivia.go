package Trivia

import (
	"math/rand"
	"time"

	Conversation "github.com/n3k0lai/ene/internal/conversation"
	Plugin "github.com/n3k0lai/ene/internal/plugins/plugin"
	Users "github.com/n3k0lai/ene/internal/users"
)

type Trivia struct {
	*Plugin.Plugin
	Questions      []Question
	ActiveQuestion Question
}

func NewTrivia(bu *Users.User) *Trivia {
	rand.Seed(time.Now().UnixNano())
	return &Trivia{
		Plugin: &Plugin.Plugin{
			Name:    "trivia",
			BotUser: *bu,
		},
		Questions: []Question{},
	}
}

func (t *Trivia) Test(query string) bool {
	if t.ActiveQuestion.Solved {
		if query == "!trivia" {
			// ask question
			t.ActiveQuestion = t.GetQuestion()
			t.ActiveQuestion.Ask()
			return true
		} else {
			return false
		}
	}

	return t.ActiveQuestion.AddQuery(query)
}

func (t *Trivia) Converse(c *Conversation.Conversation) *Conversation.Conversation {
	//add a test message to the conversation
	c.OnMessage(Conversation.NewMessage("trivia response", t.BotUser))
	return c
}

func (t *Trivia) GetQuestion() Question {
	return t.Questions[rand.Intn(len(t.Questions))]
}

func (t *Trivia) Reset() {
	t.ActiveQuestion.Reset()
}

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
	Asking         bool
}

func NewTrivia(bu *Users.User) *Trivia {
	rand.Seed(time.Now().UnixNano())
	questions := []Question{
		*NewQuestion("What is the capital of the United States?", "Washington, D.C.", 30, 0.5),
		*NewQuestion("What is the capital of Canada?", "Ottawa", 30, 0.5),
		*NewQuestion("What is the capital of Mexico?", "Mexico City", 30, 0.5),
	}
	return &Trivia{
		Plugin: &Plugin.Plugin{
			Name:    "trivia",
			BotUser: *bu,
		},
		Questions: questions,
		Asking:    false,
	}
}

func (t *Trivia) Test(c Conversation.Conversation) bool {
	if !t.Asking {
		if c.GetLatestMessage().Text == "!trivia" {
			return true
		} else {
			return false
		}
	}

	return t.ActiveQuestion.AddQuery(c.GetLatestMessage())
}

func (t *Trivia) Converse(c Conversation.Conversation) Conversation.Conversation {
	if !t.Asking {
		// ask question
		t.ActiveQuestion = t.GetQuestion()
		t.Asking = true
		questionText := t.ActiveQuestion.Ask()
		c.OnMessage(Conversation.NewMessage(questionText, t.BotUser))
	}

	if !t.ActiveQuestion.Solved {

		if t.ActiveQuestion.Expired {
			t.ActiveQuestion.Solved = true
			c.OnMessage(Conversation.NewMessage("Time's up! The answer was: "+t.ActiveQuestion.Answer, t.BotUser))
			t.Asking = false

		}
	} else {
		c.OnMessage(Conversation.NewMessage("Correct! The answer was: "+t.ActiveQuestion.Answer, t.BotUser))
		t.Asking = false
	}

	return c
}

func (t *Trivia) GetQuestion() Question {
	return t.Questions[rand.Intn(len(t.Questions))]
}

func (t *Trivia) Reset() {
	t.ActiveQuestion.Reset()
}

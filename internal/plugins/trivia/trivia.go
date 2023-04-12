package Trivia

import (
	"math/rand"
	"time"

	Lib "github.com/n3k0lai/ene/internal/lib"
)

type Trivia struct {
	Questions []Question
}

func NewTrivia(dict Lib.Dictionary) *Trivia {
	rand.Seed(time.Now().UnixNano())
	return &Trivia{
		Questions: []Question{},
	}
}

func (t *Trivia) GetQuestion() Question {
	return t.Questions[rand.Intn(len(t.Questions))]
}

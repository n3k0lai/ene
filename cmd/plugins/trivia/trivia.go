package Trivia

import "math/rand"

type Trivia struct {
	Questions []Question
}

func NewTrivia(dict Dictionary) *Trivia {
	rand.Seed(time.Now().UnixNano())
	return &Trivia{
		Questions: []Question{},
	}
}

func (t *Trivia) GetQuestion() Question {
	return t.Questions[rand.Intn(len(t.Questions))]
}

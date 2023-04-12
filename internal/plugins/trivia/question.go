package Trivia

import (
	"time"
)

type Question struct {
	Text         string
	Answer       string
	Timer        *time.Timer
	Queries      []Answer
	Solved       bool
	Limit        float64
	SecondsLimit int
	Expired      bool
	LastAsked    *time.Time
}

func NewQuestion(text string, answer string, secondsLimit int, percentageLimit float64) *Question {
	return &Question{
		Text:         text,
		Answer:       answer,
		Timer:        nil,
		Queries:      []Answer{},
		Solved:       false,
		Expired:      false,
		Limit:        percentageLimit,
		SecondsLimit: secondsLimit,
		LastAsked:    nil,
	}
}

func (q *Question) Reset() {
	q.Queries = []Answer{}
	q.Solved = false
	q.Expired = false
	q.Timer = nil
}

func (q *Question) Ask() {
	q.Reset()
	*q.LastAsked = time.Now()
	q.Timer = time.NewTimer(time.Second * time.Duration(q.SecondsLimit))
	go func() {
		<-q.Timer.C
		// hint
		if q.Solved {
			return
		}
		q.Expired = true
	}()

	// say question
}
func (q *Question) AddQuery(query string) {
	if q.Solved || q.Expired {
		return
	}
	answer := *NewAnswer(query, q.Answer)
	if answer.Percentage < q.Limit {
		q.Solved = true
		answer.Winner = true

		// notify of win

		q.Reset()
	}
	q.Queries = append(q.Queries, answer)
}

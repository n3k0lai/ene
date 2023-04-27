package Trivia

import (
	"time"

	Conversation "github.com/n3k0lai/ene/internal/conversation"
)

type Question struct {
	Text         string
	Answer       string
	Timer        time.Timer
	Queries      []Answer
	Solved       bool
	Limit        float64
	SecondsLimit int
	Expired      bool
	LastAsked    time.Time
}

func NewQuestion(text string, answer string, secondsLimit int, percentageLimit float64) *Question {
	return &Question{
		Text:         text,
		Answer:       answer,
		Timer:        time.Timer{},
		Queries:      []Answer{},
		Solved:       false,
		Expired:      false,
		Limit:        percentageLimit,
		SecondsLimit: secondsLimit,
		LastAsked:    time.Time{},
	}
}

func (q *Question) Reset() {
	q.Queries = []Answer{}
	q.Solved = false
	q.Expired = false
	q.Timer = time.Timer{}
}

func (q *Question) Ask() string {
	q.Reset()
	q.LastAsked = time.Now()
	q.Timer = *time.NewTimer(time.Second * time.Duration(q.SecondsLimit))
	// set timer for end
	go func() {
		<-q.Timer.C
		if q.Solved {
			return
		}
		q.Expired = true
	}()

	// set timer for hint
	go func() {
		<-time.After(time.Second * time.Duration(q.SecondsLimit/2))	
		if q.Solved {
			return
		}
		// send hint
	}()

	// say questionp
	return q.Text
}

func (q *Question) AddQuery(m Conversation.Message) bool {
	if q.Solved || q.Expired {
		return false
	}
	answer := *NewAnswer(m.Text, q.Answer)
	if answer.Percentage < q.Limit {
		q.Solved = true
		answer.Winner = true

		// notify of win

		q.Reset()

		return true
	}
	q.Queries = append(q.Queries, answer)

	return false
}

package Trivia

import (
	"strings"
	"time"

	Conversation "github.com/n3k0lai/ene/internal/conversation"
)

type Question struct {
	Text           string
	Answer         string
	Timer          time.Timer
	QueryStream    <-chan Conversation.Message
	ResponseStream chan<- string
	Queries        []Answer
	Solved         bool
	Limit          float64
	SecondsLimit   int
	Expired        bool
	LastAsked      time.Time
}

func NewQuestion(text string, answer string, secondsLimit int, percentageLimit float64) *Question {
	return &Question{
		Text:         text,
		Answer:       answer,
		Timer:        time.Timer{},
		QueryStream:  make(chan Conversation.Message),
		Solved:       false,
		Expired:      false,
		Limit:        percentageLimit,
		SecondsLimit: secondsLimit,
		LastAsked:    time.Time{},
	}
}

func (q *Question) Reset() {
	q.Solved = false
	q.Expired = false
	q.Timer = time.Timer{}
}
func (q *Question) Alert(msg string) {
	// send alert
	q.ResponseStream <- msg

}
func (q *Question) Ask(responseStream chan<- string) string {
	q.Reset()
	q.ResponseStream = responseStream
	q.LastAsked = time.Now()
	select {
	case query := <-q.QueryStream:
		if !q.Solved && !q.Expired {
			q.AddQuery(query)
		}
	case <-time.After(time.Duration(q.SecondsLimit/2) * time.Second):
		if !q.Solved && !q.Expired {
			// send hint
			q.Alert("Hint: " + q.Answer[0:1] + strings.Repeat("*", len(q.Answer)-1))
		}
	case <-time.After(time.Duration(q.SecondsLimit) * time.Second):
		if !q.Solved {
			q.Expired = true

			// notify of loss
			q.Alert("Time's up! The answer was: " + q.Answer)
		}
	}
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
		q.Alert("Winner: @" + m.User.Username + ": " + m.Text)

		q.Reset()

		return true
	}
	q.Queries = append(q.Queries, answer)

	return false
}

package Trivia

import Lib "github.com/n3k0lai/ene/internal/lib"

type Answer struct {
	RealAnswer string
	Query      string
	Percentage float64
	Winner     bool
}

func NewAnswer(query string, realAnswer string) *Answer {
	if query == realAnswer {
		return &Answer{
			RealAnswer: realAnswer,
			Query:      query,
			Percentage: 0.0,
			Winner:     false,
		}
	}

	return &Answer{
		RealAnswer: realAnswer,
		Query:      query,
		Percentage: Lib.GetSimilarity(query, realAnswer),
		Winner:     false,
	}
}

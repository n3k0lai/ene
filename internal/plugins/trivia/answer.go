package Trivia

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

	length := len(query)

	if len(realAnswer) < length {
		length = len(realAnswer)
	}

	differences := 0
	for i := 0; i < length; i++ {
		if query[i] != realAnswer[i] {
			differences++
		}
	}

	return &Answer{
		RealAnswer: realAnswer,
		Query:      query,
		Percentage: float64(differences) / float64(length),
		Winner:     false,
	}
}

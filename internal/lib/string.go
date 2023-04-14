package Lib

import "strings"

func GetSimilarity(a string, b string) float64 {
	if a == b {
		return 0.0
	}

	// this algo lowkey sucks, but it's good enough for now
	length := len(a)

	if len(b) < length {
		length = len(b)
	}

	differences := 0
	for i := 0; i < length; i++ {
		if a[i] != b[i] {
			differences++
		}
	}

	return float64(differences) / float64(length)
}

func CleanString(s string) string {
	//return strings.Replace(s, " ", "", -1)
	return strings.TrimSpace(s)
}

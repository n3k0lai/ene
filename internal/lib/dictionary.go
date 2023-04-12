package Lib

import (
	"bufio"
	"os"
)

type Dictionary struct {
	Lines []string
	Seed  uint64
}

func NewDictionary(filename string) *Dictionary {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return &Dictionary{
		Lines: lines,
	}
}

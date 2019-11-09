package wordcount

import (
	"strings"
)

// Frequency dict
type Freq map[string]int

func WordCount(str string) Freq {
	str = strings.Replace(str, ",", " ", -1)

	wordList := strings.Split(str, " ")
	wordCount := Freq{}

	for _, element := range wordList {
		wordCount[element]++
	}
	return wordCount
}

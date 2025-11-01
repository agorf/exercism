package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

var wordRe = regexp.MustCompile(`[a-zA-Z0-9]+('[a-zA-Z]+)?`)

func WordCount(phrase string) Frequency {
	words := wordRe.FindAllString(phrase, -1)
	freq := make(Frequency)
	for _, word := range words {
		freq[strings.ToLower(word)]++
	}
	return freq
}

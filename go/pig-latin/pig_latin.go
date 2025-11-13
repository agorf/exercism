package piglatin

import (
	"regexp"
	"strings"
)

const (
	vowels     = "aeiou"
	consonants = "bcdfghjklmnpqrstvwxyz"
)

var (
	reRule1 = regexp.MustCompile(`^([` + vowels + `]|xr|yt)`)
	reRule2 = regexp.MustCompile(`^([` + consonants + `]+)(\S*)`)
	reRule3 = regexp.MustCompile(`^([` + consonants + `]*qu)(\S*)`)
	reRule4 = regexp.MustCompile(`^([` + consonants + `]+)(y\S*)`)
)

func Sentence(sentence string) string {
	words := strings.Fields(sentence)
	var pigLatinWords []string
	for _, word := range words {
		var pigLatinWord string
		switch {
		case reRule1.MatchString(word):
			pigLatinWord = word
		case reRule3.MatchString(word):
			matches := reRule3.FindStringSubmatch(word)
			pigLatinWord = matches[2] + matches[1]
		case reRule4.MatchString(word):
			matches := reRule4.FindStringSubmatch(word)
			pigLatinWord = matches[2] + matches[1]
		case reRule2.MatchString(word):
			matches := reRule2.FindStringSubmatch(word)
			pigLatinWord = matches[2] + matches[1]
		}
		pigLatinWords = append(pigLatinWords, pigLatinWord+"ay")
	}
	return strings.Join(pigLatinWords, " ")
}

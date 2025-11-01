package anagram

import (
	"maps"
	"strings"
	"unicode/utf8"
)

func wordFreq(word string) map[rune]int {
	freq := make(map[rune]int)
	for _, r := range word {
		freq[r]++
	}
	return freq
}

func isAnagram(subject string, candidate string) bool {
	subject = strings.ToLower(subject)
	candidate = strings.ToLower(candidate)
	if subject == candidate {
		return false
	}
	if utf8.RuneCountInString(subject) != utf8.RuneCountInString(candidate) {
		return false
	}
	return maps.Equal(wordFreq(subject), wordFreq(candidate))
}

func Detect(subject string, candidates []string) []string {
	filtered := make([]string, 0)
	for _, candidate := range candidates {
		if isAnagram(subject, candidate) {
			filtered = append(filtered, candidate)
		}
	}
	return filtered
}

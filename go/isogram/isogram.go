package isogram

import "unicode"

func IsIsogram(word string) bool {
	seen := make(map[rune]bool)
	for _, r := range word {
		if r == ' ' || r == '-' {
			continue
		}
		r := unicode.ToUpper(r)
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

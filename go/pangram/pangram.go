package pangram

import "unicode"

func IsPangram(input string) bool {
	seen := make(map[rune]bool)
	total := 0
	for _, r := range input {
		r = unicode.ToLower(r)
		if r < 'a' || r > 'z' {
			continue
		}
		if seen[r] {
			continue
		}
		seen[r] = true
		total++
	}
	return total == 26
}

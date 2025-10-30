package scrabble

import "unicode"

func Score(word string) int {
	score := 0
	for _, wc := range word {
		wcu := unicode.ToUpper(wc)
		switch wcu {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			score += 1
		case 'D', 'G':
			score += 2
		case 'B', 'C', 'M', 'P':
			score += 3
		case 'F', 'H', 'V', 'W', 'Y':
			score += 4
		case 'K':
			score += 5
		case 'J', 'X':
			score += 8
		case 'Q', 'Z':
			score += 10
		}
	}
	return score
}

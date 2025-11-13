package encode

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

// Returns "" for 'X', 0
// Returns "X" for 'X', 1
// Returns "NX" for 'X', N >= 2
func encodeLetter(letter rune, count int) string {
	switch {
	case count == 1:
		return string(letter)
	case count >= 2:
		return strconv.Itoa(count) + string(letter)
	default:
		return ""
	}
}

// Returns XXX for 'X', 3
func decodeLetter(letter rune, count int) string {
	var result []rune
	for range count {
		result = append(result, letter)
	}
	return string(result)
}

func RunLengthEncode(input string) string {
	var encodedParts []string
	var rPrev rune
	n := utf8.RuneCountInString(input)
	count := 0
	i := 0
	for _, r := range input {
		if rPrev != 0 && r != rPrev {
			encodedParts = append(encodedParts, encodeLetter(rPrev, count))
			count = 0
		}
		count++
		if i == n-1 { // Last rune
			encodedParts = append(encodedParts, encodeLetter(r, count))
		}
		rPrev = r
		i++
	}
	return strings.Join(encodedParts, "")
}

func RunLengthDecode(input string) string {
	var decodedParts []string
	var countDigits []rune
	for _, r := range input {
		if r >= '0' && r <= '9' {
			countDigits = append(countDigits, r)
			continue
		}
		var count int
		var err error
		if len(countDigits) == 0 {
			count = 1
		} else {
			count, err = strconv.Atoi(string(countDigits))
			if err != nil {
				panic(err)
			}
			countDigits = []rune{}
		}
		decodedParts = append(decodedParts, decodeLetter(r, count))
	}
	return strings.Join(decodedParts, "")
}

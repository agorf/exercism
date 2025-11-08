package cryptosquare

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func normalizeMessage(plaintext string) string {
	var normalized []rune
	for _, r := range plaintext {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			r = unicode.ToLower(r)
			normalized = append(normalized, r)
		}
	}
	return string(normalized)
}

func squareDimensions(normalized string) (rows int, cols int) {
	n := len(normalized)
	s := int(math.Ceil(math.Sqrt(float64(n))))
	rows = s
	cols = s
	if (s-1)*s >= n {
		rows--
	}
	return rows, cols
}

func Encode(pt string) string {
	normalized := normalizeMessage(pt)
	rows, cols := squareDimensions(normalized)
	n := len(normalized)
	chunks := make([]string, 0, cols)
	for c := range cols {
		var chunk []byte
		for r := range rows {
			i := r*cols + c
			if i < n {
				chunk = append(chunk, normalized[i])
			}
		}
		chunkStr := fmt.Sprintf("%-*s", rows, string(chunk)) // Right-pad
		chunks = append(chunks, chunkStr)
	}
	return strings.Join(chunks, " ")
}

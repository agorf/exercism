package atbash

import "unicode"

func Atbash(s string) string {
	var cipher []rune
	count := 0
	appendGrouped := func(r rune) {
		if count == 5 {
			cipher = append(cipher, ' ')
			count = 0
		}
		cipher = append(cipher, r)
		count++
	}
	for _, r := range s {
		if r >= '0' && r <= '9' {
			appendGrouped(r)
			continue
		}
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			r = unicode.ToLower(r)
			appendGrouped('z' - (r - 'a'))
		}
	}
	return string(cipher)
}

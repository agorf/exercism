package rotationalcipher

import (
	"unicode"
	"unicode/utf8"
)

func ShiftLetter(r rune, shiftKey int) rune {
	a := 'a'
	if unicode.IsUpper(r) {
		a = 'A'
	}
	return rune(((int(r-a) + shiftKey) % 26)) + a
}

func RotationalCipher(plain string, shiftKey int) string {
	cipher := make([]rune, utf8.RuneCountInString(plain))
	ri := 0
	for _, r := range plain {
		c := r
		if unicode.IsLetter(r) {
			c = ShiftLetter(r, shiftKey)
		}
		cipher[ri] = c
		ri++
	}
	return string(cipher)
}

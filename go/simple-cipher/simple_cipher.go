package cipher

import (
	"regexp"
	"strings"
	"unicode"
)

type (
	shift struct {
		distance int
	}
	vigenere struct {
		key string
	}
)

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance < -25 || distance == 0 || distance > 25 {
		return nil
	}
	return shift{distance}
}

func (c shift) crypt(input string, multiplier int) string {
	var out []rune
	for _, r := range input {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			r = unicode.ToLower(r)
			c := 'a' + rune((((int(r-'a')+(multiplier*c.distance))%26)+26)%26)
			out = append(out, c)
		}
	}
	return string(out)
}

func (c shift) Encode(input string) string {
	return c.crypt(input, 1)
}

func (c shift) Decode(input string) string {
	return c.crypt(input, -1)
}

func NewVigenere(key string) Cipher {
	if strings.TrimRight(key, "a") == "" { // "", "a", "aa" etc
		return nil
	}
	if regexp.MustCompile(`[^a-z]`).MatchString(key) { // Anything other than a-z
		return nil
	}
	return vigenere{key}
}

func (v vigenere) crypt(input string, multiplier int) string {
	var out []rune
	keyIndex := 0
	for _, r := range input {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			r = unicode.ToLower(r)
			distance := int(v.key[keyIndex%len(v.key)] - 'a')
			c := 'a' + rune((((int(r-'a')+(multiplier*distance))%26)+26)%26)
			out = append(out, c)
			keyIndex++
		}
	}
	return string(out)
}

func (v vigenere) Encode(input string) string {
	return v.crypt(input, 1)
}

func (v vigenere) Decode(input string) string {
	return v.crypt(input, -1)
}

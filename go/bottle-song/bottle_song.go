package bottlesong

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

var numbers = [...]string{
	"no",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
}

func pluralize(singular, plural string, n int) string {
	if n == 1 {
		return singular
	}
	return plural
}

func titleize(s string) string {
	r, size := utf8.DecodeRuneInString(s)
	r = unicode.ToUpper(r)
	return string(r) + s[size:]
}

func hangingLine(bottles int) string {
	return fmt.Sprintf(
		"%s green %s hanging on the wall,",
		titleize(numbers[bottles]),
		pluralize("bottle", "bottles", bottles),
	)
}

func finalLine(bottles int) string {
	return fmt.Sprintf(
		"There'll be %s green %s hanging on the wall.",
		numbers[bottles],
		pluralize("bottle", "bottles", bottles),
	)
}

func Recite(startBottles, takeDown int) []string {
	var song []string
	bottles := startBottles
	for i := takeDown; i > 0; i-- {
		line := hangingLine(bottles)
		song = append(song, line)
		song = append(song, line)
		song = append(song, "And if one green bottle should accidentally fall,")
		song = append(song, finalLine(bottles-1))
		if i-1 > 0 {
			song = append(song, "")
		}
		bottles--
	}
	return song
}

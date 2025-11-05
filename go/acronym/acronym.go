package acronym

import "unicode"

func Abbreviate(s string) string {
	var acronym []rune
	capture := true
	for _, r := range s {
		if r == ' ' || r == '-' {
			capture = true
			continue
		}
		if capture && unicode.IsLetter(r) {
			acronym = append(acronym, unicode.ToUpper(r))
			capture = false
		}
	}
	return string(acronym)
}

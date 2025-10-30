package reverse

import "unicode/utf8"

func Reverse(input string) string {
	length := utf8.RuneCountInString(input)
	reverse := make([]rune, length)
	i := 0
	for _, r := range input {
		reverse[length-1-i] = r
		i++
	}
	return string(reverse)
}

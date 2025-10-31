package bob

import "regexp"

func MustMatchString(pattern string, s string) bool {
	matched, err := regexp.MatchString(pattern, s)
	if err != nil {
		panic(err)
	}
	return matched
}

func Question(remark string) bool {
	return MustMatchString(`\?\s*$`, remark)
}

func Yelling(remark string) bool {
	lower := MustMatchString(`[a-z]`, remark)
	upper := MustMatchString(`[A-Z]`, remark)
	return upper && !lower
}

func Silence(remark string) bool {
	return MustMatchString(`^\s*$`, remark)
}

func Hey(remark string) (answer string) {
	question := Question(remark)
	yelling := Yelling(remark)
	silence := Silence(remark)
	switch {
	case question && !yelling:
		answer = "Sure."
	case !question && yelling:
		answer = "Whoa, chill out!"
	case question && yelling:
		answer = "Calm down, I know what I'm doing!"
	case silence:
		answer = "Fine. Be that way!"
	default:
		answer = "Whatever."
	}
	return answer
}

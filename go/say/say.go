package say

import "strings"

type divisor struct {
	num  int64
	word string
}

var divisors = []divisor{
	{1000000000, "billion"},
	{1000000, "million"},
	{1000, "thousand"},
	{100, "hundred"},
	{1, ""},
}

var numWords = map[int64]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

func SayHundred(n int64) ([]string, bool) {
	if n >= 1000 {
		return []string{}, false
	}
	var parts []string
	if n >= 100 {
		parts = append(parts, numWords[n/100]) // 987 / 100 == 9
		parts = append(parts, "hundred")
		n %= 100 // 987 % 100 = 87
	}
	if n >= 10 {
		word, exists := numWords[n]
		if exists {
			parts = append(parts, word)
			return parts, true
		}
		parts = append(parts, numWords[(n/10)*10]) // (87/10)*10 == (8)*10 == 80
		n %= 10                                    // 87 % 10 == 7
	}
	if n > 0 {
		parts = append(parts, numWords[n])
	}
	nParts := len(parts)
	if nParts >= 2 {
		last := parts[nParts-2] + "-" + parts[nParts-1] // Join last two with dash
		parts = parts[:nParts-2]                        // Drop last two
		parts = append(parts, last)                     // Add back last two joined with dash
	}
	return parts, true
}

func Say(n int64) (string, bool) {
	if n < 0 {
		return "", false
	}
	if n == 0 {
		return "zero", true
	}
	var parts []string
	for _, div := range divisors {
		if n < div.num {
			continue
		}
		hundredParts, ok := SayHundred(n / div.num)
		if !ok {
			return "", false
		}
		parts = append(parts, hundredParts...)
		if div.word != "" {
			parts = append(parts, div.word)
		}
		n %= div.num
	}
	return strings.Join(parts, " "), true
}

package romannumerals

import (
	"fmt"
	"strings"
)

const (
	minInput = 1
	maxInput = 3999
)

type numeralPair struct {
	arabic int
	roman  string
}

var arabicToRoman = [...]numeralPair{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

var errOutOfRange = fmt.Errorf("input should be between %d and %d", minInput, maxInput)

func ToRomanNumeral(input int) (string, error) {
	if input < minInput || input > maxInput {
		return "", errOutOfRange
	}

	roman := ""
	for _, pair := range arabicToRoman {
		count := input / pair.arabic
		if count == 0 {
			continue
		}

		roman += strings.Repeat(pair.roman, count)

		input %= pair.arabic
		if input == 0 {
			break
		}
	}
	return roman, nil
}

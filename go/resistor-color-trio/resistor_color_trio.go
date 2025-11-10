package resistorcolortrio

import "fmt"

var colorsToNumbers = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

func intPow(base, exp int) int {
	if exp == 0 {
		return 1
	}
	result := 1
	for exp > 0 {
		result *= base
		exp--
	}
	return result
}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	zeros := colorsToNumbers[colors[2]]
	value := colorsToNumbers[colors[0]]*intPow(10, zeros+1) + colorsToNumbers[colors[1]]*intPow(10, zeros)
	var unitPrefix string
	switch {
	case value >= 1000000000:
		unitPrefix = "giga"
		value /= 1000000000
	case value >= 1000000:
		unitPrefix = "mega"
		value /= 1000000
	case value >= 1000:
		unitPrefix = "kilo"
		value /= 1000
	}
	return fmt.Sprintf("%d %sohms", value, unitPrefix)
}

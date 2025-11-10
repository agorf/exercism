package allyourbase

import (
	"errors"
)

func intPow(base, exp int) int {
	if exp == 0 {
		return 1
	}
	result := base
	for range exp - 1 {
		result *= base
	}
	return result
}

func divMod(dividend, divisor int) (int, int) {
	div := dividend / divisor
	mod := dividend % divisor
	return div, mod
}

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}
	for _, digit := range inputDigits {
		if digit < 0 || digit >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
	}
	if inputBase == outputBase {
		return inputDigits, nil
	}
	inputValue := 0
	n := len(inputDigits)
	for i, digit := range inputDigits {
		inputValue += digit * intPow(inputBase, n-i-1)
	}
	if inputValue == 0 {
		return []int{0}, nil
	}
	var outputDigits []int
	for inputValue > 0 {
		var digit int
		inputValue, digit = divMod(inputValue, outputBase)
		outputDigits = append([]int{digit}, outputDigits...) // Prepend
	}
	return outputDigits, nil
}

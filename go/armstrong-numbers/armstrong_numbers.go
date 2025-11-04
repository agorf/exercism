package armstrong

func intPow(base, exponent int) int {
	result := 1
	for range exponent {
		result *= base
	}
	return result
}

func IsNumber(n int) bool {
	sum := 0
	pow := 0
	for t := n; t > 0; t /= 10 { // For n == 154, t = 154, 15, 1, 0
		pow++
	}
	for t := n; t > 0; t /= 10 {
		d := t % 10
		sum += intPow(d, pow)
	}
	return sum == n
}

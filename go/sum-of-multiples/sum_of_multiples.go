package summultiples

func SumMultiples(limit int, divisors ...int) int {
	multiples := make(map[int]struct{})
	for _, divisor := range divisors {
		if divisor == 0 {
			continue
		}
		for i := divisor; i < limit; i += divisor {
			multiples[i] = struct{}{}
		}
	}
	sum := 0
	for multiple := range multiples {
		sum += multiple
	}
	return sum
}

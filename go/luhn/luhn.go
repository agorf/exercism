package luhn

func Valid(id string) bool {
	double := false
	sum := 0
	length := 0
	for i := len(id) - 1; i >= 0; i-- {
		b := id[i]
		if b == ' ' {
			continue
		}
		if b < '0' || b > '9' {
			return false
		}
		n := int(b - '0') // Convert '0' to 0, '1' to 1, etc
		if double {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		double = !double
		sum += n
		length++
	}
	return length > 1 && sum%10 == 0
}

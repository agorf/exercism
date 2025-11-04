package series

func All(n int, s string) []string {
	var series []string
	for i := 0; i+n <= len(s); i++ {
		series = append(series, s[i:i+n])
	}
	return series
}

func UnsafeFirst(n int, s string) string {
	return All(n, s)[0]
}

func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "", false
	}
	return UnsafeFirst(n, s), true
}

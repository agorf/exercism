package rectangles

type cols struct {
	from, till int
}

func isSquareLine(line string, c cols) bool {
	if line[c.from] != '+' || line[c.till] != '+' {
		return false
	}
	for i := c.from + 1; i < c.till; i++ {
		if line[i] != '+' && line[i] != '-' {
			return false
		}
	}
	return true
}

func isSquareSides(line string, c cols) bool {
	return (line[c.from] == '+' || line[c.from] == '|') && (line[c.till] == '+' || line[c.till] == '|')
}

func Count(diagram []string) int {
	squares := make(map[cols]int)
	count := 0
	for _, line := range diagram {
		n := len(line)
		for i := 0; i < n-1; i++ {
			for j := i + 1; j < n; j++ {
				c := cols{i, j}
				times, seen := squares[c]
				if isSquareLine(line, c) {
					if seen {
						count += times
					}
					squares[c]++
				} else if !isSquareSides(line, c) {
					delete(squares, c)
				}
			}
		}
	}
	return count
}

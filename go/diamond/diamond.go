package diamond

import (
	"errors"
	"strings"
)

var errOutOfRange = errors.New("character must be a capital letter")

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", errOutOfRange
	}
	size := (int(char-'A')+1)*2 - 1
	mid := size / 2
	var result []string
	for i := range size {
		row := make([]byte, size)
		for j := range size {
			row[j] = ' '
		}
		var x int
		if i <= mid {
			x = mid - i
		} else if i > mid {
			x = i - mid
		}
		letter := char - byte(x)
		row[x] = letter
		row[size-x-1] = letter
		result = append(result, string(row))
	}
	return strings.Join(result, "\n"), nil
}

package wordsearch

import "errors"

var errNotFound = errors.New("not found")

func collectWordTopRight(puzzle []string, from, till [2]int) string {
	var chars []byte
	for i := from[1]; i >= till[1]; i-- {
		for j := from[0]; j <= till[0]; j++ {
			if from[1]-i == j-from[0] { // Diagonal
				chars = append(chars, puzzle[i][j])
			}
		}
	}
	return string(chars)
}

func collectWordRight(puzzle []string, from, till [2]int) string {
	var chars []byte
	i := from[1]
	for j := from[0]; j <= till[0]; j++ {
		chars = append(chars, puzzle[i][j])
	}
	return string(chars)
}

func collectWordBottomRight(puzzle []string, from, till [2]int) string {
	var chars []byte
	for i := from[1]; i <= till[1]; i++ {
		for j := from[0]; j <= till[0]; j++ {
			if i-from[1] == j-from[0] { // Diagonal
				chars = append(chars, puzzle[i][j])
			}
		}
	}
	return string(chars)
}

func collectWordBottom(puzzle []string, from, till [2]int) string {
	var chars []byte
	j := from[0]
	for i := from[1]; i <= till[1]; i++ {
		chars = append(chars, puzzle[i][j])
	}
	return string(chars)
}

func collectWordBottomLeft(puzzle []string, from, till [2]int) string {
	var chars []byte
	for i := from[1]; i <= till[1]; i++ {
		for j := from[0]; j >= till[0]; j-- {
			if i-from[1] == from[0]-j { // Diagonal
				chars = append(chars, puzzle[i][j])
			}
		}
	}
	return string(chars)
}

func collectWordLeft(puzzle []string, from, till [2]int) string {
	var chars []byte
	i := from[1]
	for j := from[0]; j >= till[0]; j-- {
		chars = append(chars, puzzle[i][j])
	}
	return string(chars)
}

func collectWordTopLeft(puzzle []string, from, till [2]int) string {
	var chars []byte
	for i := from[1]; i >= till[1]; i-- {
		for j := from[0]; j >= till[0]; j-- {
			if from[1]-i == from[0]-j { // Diagonal
				chars = append(chars, puzzle[i][j])
			}
		}
	}
	return string(chars)
}

func collectWordTop(puzzle []string, from, till [2]int) string {
	var chars []byte
	j := from[0]
	for i := from[1]; i >= till[1]; i-- {
		chars = append(chars, puzzle[i][j])
	}
	return string(chars)
}

func sanitizeIndex(value, min, max int) int {
	if value < min || value > max {
		return -1
	}
	return value
}

func wordPosition(puzzle []string, word string, from [2]int) ([2]int, error) {
	wordLen := len(word)
	maxRow := len(puzzle) - 1
	maxCol := len(puzzle[0]) - 1
	top := sanitizeIndex(from[1]-wordLen+1, 0, maxRow)
	right := sanitizeIndex(from[0]+wordLen-1, 0, maxCol)
	bottom := sanitizeIndex(from[1]+wordLen-1, 0, maxRow)
	left := sanitizeIndex(from[0]-wordLen+1, 0, maxCol)
	if top >= 0 {
		till := [2]int{from[0], top}
		if collectWordTop(puzzle, from, till) == word {
			return till, nil
		}
	}
	if top >= 0 && right >= 0 {
		till := [2]int{right, top}
		if collectWordTopRight(puzzle, from, till) == word {
			return till, nil
		}
	}
	if right >= 0 {
		till := [2]int{right, from[1]}
		if collectWordRight(puzzle, from, till) == word {
			return till, nil
		}
	}
	if bottom >= 0 && right >= 0 {
		till := [2]int{right, bottom}
		if collectWordBottomRight(puzzle, from, till) == word {
			return till, nil
		}
	}
	if bottom >= 0 {
		till := [2]int{from[0], bottom}
		if collectWordBottom(puzzle, from, till) == word {
			return till, nil
		}
	}
	if bottom >= 0 && left >= 0 {
		till := [2]int{left, bottom}
		if collectWordBottomLeft(puzzle, from, till) == word {
			return till, nil
		}
	}
	if left >= 0 {
		till := [2]int{left, from[1]}
		if collectWordLeft(puzzle, from, till) == word {
			return till, nil
		}
	}
	if top >= 0 && left >= 0 {
		till := [2]int{left, top}
		if collectWordTopLeft(puzzle, from, till) == word {
			return till, nil
		}
	}
	return [2]int{}, errNotFound
}

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	results := map[string][2][2]int{}
	numRows := len(puzzle)
	numCols := len(puzzle[0])
	for _, word := range words {
		for i := range numRows {
			for j := range numCols {
				from := [2]int{j, i}
				till, err := wordPosition(puzzle, word, from)
				if err == nil { // Found
					results[word] = [2][2]int{from, till}
				}
			}
		}
		if _, wordFound := results[word]; !wordFound {
			return results, errNotFound
		}
	}
	return results, nil
}

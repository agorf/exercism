package matrix

import "strings"

type Matrix [][]int

type Pair struct {
	row, col int
}

func New(s string) (*Matrix, error) {
	var m Matrix
	for _, line := range strings.Split(s, "\n") {
		var row []int
		if line != "" {
			for _, numStr := range strings.Split(line, " ") {
				row = append(row, int(numStr[0]-'0'))
			}
		}
		m = append(m, row)
	}
	return &m, nil
}

func (m *Matrix) Saddle() []Pair {
	var pairs []Pair
	numRows := len(*m)
	numCols := len((*m)[0])
	maxRows := make([]int, numRows)
	minCols := make([]int, numCols)
	for j := range minCols {
		minCols[j] = 9 // So that min later works
	}
	for i := range numRows {
		for j := range numCols {
			num := (*m)[i][j]
			maxRows[i] = max(maxRows[i], num)
			minCols[j] = min(minCols[j], num)
		}
	}
	for i, r := range maxRows {
		for j, c := range minCols {
			if r == c {
				pairs = append(pairs, Pair{i + 1, j + 1})
			}
		}
	}
	return pairs
}

package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(s string) (Matrix, error) {
	var m Matrix
	var lastRow []int
	for _, row := range strings.Split(s, "\n") {
		var mRow []int
		for _, col := range strings.Fields(row) {
			num, err := strconv.Atoi(col)
			if err != nil {
				return nil, err
			}
			mRow = append(mRow, num)
		}
		if len(mRow) == 0 {
			return nil, errors.New("empty row")
		}
		if lastRow != nil && len(lastRow) != len(mRow) {
			return nil, errors.New("uneven rows")
		}
		m = append(m, mRow)
		lastRow = mRow
	}
	return m, nil
}

func (m Matrix) Dimensions() (rows, cols int) {
	return len(m), len(m[0])
}

func (m Matrix) Cols() [][]int {
	numRows, numCols := m.Dimensions()
	cols := make([][]int, numCols)
	for j := range numCols {
		cols[j] = make([]int, 0, numRows) // Optional
		for i := range numRows {
			cols[j] = append(cols[j], m[i][j])
		}
	}
	return cols
}

func (m Matrix) Rows() [][]int {
	rows := make([][]int, 0, len(m))
	for _, row := range m {
		rowCopy := append([]int{}, row...)
		rows = append(rows, rowCopy)
	}
	return rows
}

func (m Matrix) Set(row, col, val int) bool {
	numRows, numCols := m.Dimensions()
	if row < 0 || row > numRows-1 || col < 0 || col > numCols-1 {
		return false
	}
	m[row][col] = val
	return true
}
